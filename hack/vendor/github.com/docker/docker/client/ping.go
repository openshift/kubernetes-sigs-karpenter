/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"context"
	"net/http"
	"path"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/build"
	"github.com/docker/docker/api/types/swarm"
)

// Ping pings the server and returns the value of the "Docker-Experimental",
// "Builder-Version", "OS-Type" & "API-Version" headers. It attempts to use
// a HEAD request on the endpoint, but falls back to GET if HEAD is not supported
// by the daemon. It ignores internal server errors returned by the API, which
// may be returned if the daemon is in an unhealthy state, but returns errors
// for other non-success status codes, failing to connect to the API, or failing
// to parse the API response.
func (cli *Client) Ping(ctx context.Context) (types.Ping, error) {
	var ping types.Ping

	// Using cli.buildRequest() + cli.doRequest() instead of cli.sendRequest()
	// because ping requests are used during API version negotiation, so we want
	// to hit the non-versioned /_ping endpoint, not /v1.xx/_ping
	req, err := cli.buildRequest(ctx, http.MethodHead, path.Join(cli.basePath, "/_ping"), nil, nil)
	if err != nil {
		return ping, err
	}
	resp, err := cli.doRequest(req)
	if err != nil {
		if IsErrConnectionFailed(err) {
			return ping, err
		}
		// We managed to connect, but got some error; continue and try GET request.
	} else {
		defer ensureReaderClosed(resp)
		switch resp.StatusCode {
		case http.StatusOK, http.StatusInternalServerError:
			// Server handled the request, so parse the response
			return parsePingResponse(cli, resp)
		}
	}

	// HEAD failed; fallback to GET.
	req.Method = http.MethodGet
	resp, err = cli.doRequest(req)
	defer ensureReaderClosed(resp)
	if err != nil {
		return ping, err
	}
	return parsePingResponse(cli, resp)
}

func parsePingResponse(cli *Client, resp *http.Response) (types.Ping, error) {
	if resp == nil {
		return types.Ping{}, nil
	}

	var ping types.Ping
	if resp.Header == nil {
		return ping, cli.checkResponseErr(resp)
	}
	ping.APIVersion = resp.Header.Get("Api-Version")
	ping.OSType = resp.Header.Get("Ostype")
	if resp.Header.Get("Docker-Experimental") == "true" {
		ping.Experimental = true
	}
	if bv := resp.Header.Get("Builder-Version"); bv != "" {
		ping.BuilderVersion = build.BuilderVersion(bv)
	}
	if si := resp.Header.Get("Swarm"); si != "" {
		state, role, _ := strings.Cut(si, "/")
		ping.SwarmStatus = &swarm.Status{
			NodeState:        swarm.LocalNodeState(state),
			ControlAvailable: role == "manager",
		}
	}
	return ping, cli.checkResponseErr(resp)
}

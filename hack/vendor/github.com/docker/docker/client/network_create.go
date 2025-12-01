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
	"encoding/json"

	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/versions"
)

// NetworkCreate creates a new network in the docker host.
func (cli *Client) NetworkCreate(ctx context.Context, name string, options network.CreateOptions) (network.CreateResponse, error) {
	// Make sure we negotiated (if the client is configured to do so),
	// as code below contains API-version specific handling of options.
	//
	// Normally, version-negotiation (if enabled) would not happen until
	// the API request is made.
	if err := cli.checkVersion(ctx); err != nil {
		return network.CreateResponse{}, err
	}

	networkCreateRequest := network.CreateRequest{
		CreateOptions: options,
		Name:          name,
	}
	if versions.LessThan(cli.version, "1.44") {
		enabled := true
		networkCreateRequest.CheckDuplicate = &enabled //nolint:staticcheck // ignore SA1019: CheckDuplicate is deprecated since API v1.44.
	}

	resp, err := cli.post(ctx, "/networks/create", nil, networkCreateRequest, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return network.CreateResponse{}, err
	}

	var response network.CreateResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

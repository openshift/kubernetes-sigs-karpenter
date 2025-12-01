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
	"net/url"
	"strconv"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/versions"
)

// ContainerStop stops a container. In case the container fails to stop
// gracefully within a time frame specified by the timeout argument,
// it is forcefully terminated (killed).
//
// If the timeout is nil, the container's StopTimeout value is used, if set,
// otherwise the engine default. A negative timeout value can be specified,
// meaning no timeout, i.e. no forceful termination is performed.
func (cli *Client) ContainerStop(ctx context.Context, containerID string, options container.StopOptions) error {
	containerID, err := trimID("container", containerID)
	if err != nil {
		return err
	}

	query := url.Values{}
	if options.Timeout != nil {
		query.Set("t", strconv.Itoa(*options.Timeout))
	}
	if options.Signal != "" {
		// Make sure we negotiated (if the client is configured to do so),
		// as code below contains API-version specific handling of options.
		//
		// Normally, version-negotiation (if enabled) would not happen until
		// the API request is made.
		if err := cli.checkVersion(ctx); err != nil {
			return err
		}
		if versions.GreaterThanOrEqualTo(cli.version, "1.42") {
			query.Set("signal", options.Signal)
		}
	}
	resp, err := cli.post(ctx, "/containers/"+containerID+"/stop", query, nil, nil)
	ensureReaderClosed(resp)
	return err
}

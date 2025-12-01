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
	"net/url"
	"strings"

	"github.com/docker/docker/api/types/container"
)

// ContainerTop shows process information from within a container.
func (cli *Client) ContainerTop(ctx context.Context, containerID string, arguments []string) (container.TopResponse, error) {
	containerID, err := trimID("container", containerID)
	if err != nil {
		return container.TopResponse{}, err
	}

	query := url.Values{}
	if len(arguments) > 0 {
		query.Set("ps_args", strings.Join(arguments, " "))
	}

	resp, err := cli.get(ctx, "/containers/"+containerID+"/top", query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return container.TopResponse{}, err
	}

	var response container.TopResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

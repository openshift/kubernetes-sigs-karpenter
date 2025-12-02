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

	"github.com/docker/docker/api/types/container"
)

// ContainerDiff shows differences in a container filesystem since it was started.
func (cli *Client) ContainerDiff(ctx context.Context, containerID string) ([]container.FilesystemChange, error) {
	containerID, err := trimID("container", containerID)
	if err != nil {
		return nil, err
	}

	resp, err := cli.get(ctx, "/containers/"+containerID+"/changes", url.Values{}, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return nil, err
	}

	var changes []container.FilesystemChange
	err = json.NewDecoder(resp.Body).Decode(&changes)
	if err != nil {
		return nil, err
	}
	return changes, err
}

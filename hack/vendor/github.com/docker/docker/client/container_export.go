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
	"io"
	"net/url"
)

// ContainerExport retrieves the raw contents of a container
// and returns them as an io.ReadCloser. It's up to the caller
// to close the stream.
func (cli *Client) ContainerExport(ctx context.Context, containerID string) (io.ReadCloser, error) {
	containerID, err := trimID("container", containerID)
	if err != nil {
		return nil, err
	}

	resp, err := cli.get(ctx, "/containers/"+containerID+"/export", url.Values{}, nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

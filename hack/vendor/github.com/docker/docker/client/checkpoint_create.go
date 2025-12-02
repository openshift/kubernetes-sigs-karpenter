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

	"github.com/docker/docker/api/types/checkpoint"
)

// CheckpointCreate creates a checkpoint from the given container with the given name
func (cli *Client) CheckpointCreate(ctx context.Context, containerID string, options checkpoint.CreateOptions) error {
	containerID, err := trimID("container", containerID)
	if err != nil {
		return err
	}

	resp, err := cli.post(ctx, "/containers/"+containerID+"/checkpoints", nil, options, nil)
	ensureReaderClosed(resp)
	return err
}

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
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/docker/docker/api/types/swarm"
)

// TaskInspectWithRaw returns the task information and its raw representation.
func (cli *Client) TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error) {
	taskID, err := trimID("task", taskID)
	if err != nil {
		return swarm.Task{}, nil, err
	}

	resp, err := cli.get(ctx, "/tasks/"+taskID, nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return swarm.Task{}, nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return swarm.Task{}, nil, err
	}

	var response swarm.Task
	rdr := bytes.NewReader(body)
	err = json.NewDecoder(rdr).Decode(&response)
	return response, body, err
}

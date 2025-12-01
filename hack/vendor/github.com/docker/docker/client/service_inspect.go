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
	"fmt"
	"io"
	"net/url"

	"github.com/docker/docker/api/types/swarm"
)

// ServiceInspectWithRaw returns the service information and the raw data.
func (cli *Client) ServiceInspectWithRaw(ctx context.Context, serviceID string, opts swarm.ServiceInspectOptions) (swarm.Service, []byte, error) {
	serviceID, err := trimID("service", serviceID)
	if err != nil {
		return swarm.Service{}, nil, err
	}

	query := url.Values{}
	query.Set("insertDefaults", fmt.Sprintf("%v", opts.InsertDefaults))
	resp, err := cli.get(ctx, "/services/"+serviceID, query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return swarm.Service{}, nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return swarm.Service{}, nil, err
	}

	var response swarm.Service
	rdr := bytes.NewReader(body)
	err = json.NewDecoder(rdr).Decode(&response)
	return response, body, err
}

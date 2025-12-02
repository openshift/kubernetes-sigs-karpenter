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
	"fmt"
	"net/url"

	"github.com/docker/docker/api/types"
)

// DiskUsage requests the current data usage from the daemon
func (cli *Client) DiskUsage(ctx context.Context, options types.DiskUsageOptions) (types.DiskUsage, error) {
	var query url.Values
	if len(options.Types) > 0 {
		query = url.Values{}
		for _, t := range options.Types {
			query.Add("type", string(t))
		}
	}

	resp, err := cli.get(ctx, "/system/df", query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return types.DiskUsage{}, err
	}

	var du types.DiskUsage
	if err := json.NewDecoder(resp.Body).Decode(&du); err != nil {
		return types.DiskUsage{}, fmt.Errorf("Error retrieving disk usage: %v", err)
	}
	return du, nil
}

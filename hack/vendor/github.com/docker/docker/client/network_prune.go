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

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
)

// NetworksPrune requests the daemon to delete unused networks
func (cli *Client) NetworksPrune(ctx context.Context, pruneFilters filters.Args) (network.PruneReport, error) {
	if err := cli.NewVersionError(ctx, "1.25", "network prune"); err != nil {
		return network.PruneReport{}, err
	}

	query, err := getFiltersQuery(pruneFilters)
	if err != nil {
		return network.PruneReport{}, err
	}

	resp, err := cli.post(ctx, "/networks/prune", query, nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return network.PruneReport{}, err
	}

	var report network.PruneReport
	if err := json.NewDecoder(resp.Body).Decode(&report); err != nil {
		return network.PruneReport{}, fmt.Errorf("Error retrieving network prune report: %v", err)
	}

	return report, nil
}

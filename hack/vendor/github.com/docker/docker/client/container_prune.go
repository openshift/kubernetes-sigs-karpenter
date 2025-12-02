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

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
)

// ContainersPrune requests the daemon to delete unused data
func (cli *Client) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (container.PruneReport, error) {
	if err := cli.NewVersionError(ctx, "1.25", "container prune"); err != nil {
		return container.PruneReport{}, err
	}

	query, err := getFiltersQuery(pruneFilters)
	if err != nil {
		return container.PruneReport{}, err
	}

	resp, err := cli.post(ctx, "/containers/prune", query, nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return container.PruneReport{}, err
	}

	var report container.PruneReport
	if err := json.NewDecoder(resp.Body).Decode(&report); err != nil {
		return container.PruneReport{}, fmt.Errorf("Error retrieving disk usage: %v", err)
	}

	return report, nil
}

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
	"github.com/docker/docker/api/types/volume"
)

// VolumesPrune requests the daemon to delete unused data
func (cli *Client) VolumesPrune(ctx context.Context, pruneFilters filters.Args) (volume.PruneReport, error) {
	if err := cli.NewVersionError(ctx, "1.25", "volume prune"); err != nil {
		return volume.PruneReport{}, err
	}

	query, err := getFiltersQuery(pruneFilters)
	if err != nil {
		return volume.PruneReport{}, err
	}

	resp, err := cli.post(ctx, "/volumes/prune", query, nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return volume.PruneReport{}, err
	}

	var report volume.PruneReport
	if err := json.NewDecoder(resp.Body).Decode(&report); err != nil {
		return volume.PruneReport{}, fmt.Errorf("Error retrieving volume prune report: %v", err)
	}

	return report, nil
}

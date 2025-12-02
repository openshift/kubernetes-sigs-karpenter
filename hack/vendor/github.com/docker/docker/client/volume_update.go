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
	"net/url"

	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/api/types/volume"
)

// VolumeUpdate updates a volume. This only works for Cluster Volumes, and
// only some fields can be updated.
func (cli *Client) VolumeUpdate(ctx context.Context, volumeID string, version swarm.Version, options volume.UpdateOptions) error {
	volumeID, err := trimID("volume", volumeID)
	if err != nil {
		return err
	}
	if err := cli.NewVersionError(ctx, "1.42", "volume update"); err != nil {
		return err
	}

	query := url.Values{}
	query.Set("version", version.String())

	resp, err := cli.put(ctx, "/volumes/"+volumeID, query, options, nil)
	ensureReaderClosed(resp)
	return err
}

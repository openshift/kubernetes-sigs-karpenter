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

	"github.com/docker/docker/api/types/versions"
)

// VolumeRemove removes a volume from the docker host.
func (cli *Client) VolumeRemove(ctx context.Context, volumeID string, force bool) error {
	volumeID, err := trimID("volume", volumeID)
	if err != nil {
		return err
	}

	query := url.Values{}
	if force {
		// Make sure we negotiated (if the client is configured to do so),
		// as code below contains API-version specific handling of options.
		//
		// Normally, version-negotiation (if enabled) would not happen until
		// the API request is made.
		if err := cli.checkVersion(ctx); err != nil {
			return err
		}
		if versions.GreaterThanOrEqualTo(cli.version, "1.25") {
			query.Set("force", "1")
		}
	}
	resp, err := cli.delete(ctx, "/volumes/"+volumeID, query, nil)
	defer ensureReaderClosed(resp)
	return err
}

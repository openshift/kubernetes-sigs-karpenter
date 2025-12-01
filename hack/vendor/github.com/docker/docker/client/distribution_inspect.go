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
	"net/http"
	"net/url"

	"github.com/docker/docker/api/types/registry"
)

// DistributionInspect returns the image digest with the full manifest.
func (cli *Client) DistributionInspect(ctx context.Context, imageRef, encodedRegistryAuth string) (registry.DistributionInspect, error) {
	if imageRef == "" {
		return registry.DistributionInspect{}, objectNotFoundError{object: "distribution", id: imageRef}
	}

	if err := cli.NewVersionError(ctx, "1.30", "distribution inspect"); err != nil {
		return registry.DistributionInspect{}, err
	}

	var headers http.Header
	if encodedRegistryAuth != "" {
		headers = http.Header{
			registry.AuthHeader: {encodedRegistryAuth},
		}
	}

	// Contact the registry to retrieve digest and platform information
	resp, err := cli.get(ctx, "/distribution/"+imageRef+"/json", url.Values{}, headers)
	defer ensureReaderClosed(resp)
	if err != nil {
		return registry.DistributionInspect{}, err
	}

	var distributionInspect registry.DistributionInspect
	err = json.NewDecoder(resp.Body).Decode(&distributionInspect)
	return distributionInspect, err
}

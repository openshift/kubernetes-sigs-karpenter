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
	"net/url"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/versions"
)

// ImageList returns a list of images in the docker host.
//
// Experimental: Setting the [options.Manifest] will populate
// [image.Summary.Manifests] with information about image manifests.
// This is experimental and might change in the future without any backward
// compatibility.
func (cli *Client) ImageList(ctx context.Context, options image.ListOptions) ([]image.Summary, error) {
	var images []image.Summary

	// Make sure we negotiated (if the client is configured to do so),
	// as code below contains API-version specific handling of options.
	//
	// Normally, version-negotiation (if enabled) would not happen until
	// the API request is made.
	if err := cli.checkVersion(ctx); err != nil {
		return images, err
	}

	query := url.Values{}

	optionFilters := options.Filters
	referenceFilters := optionFilters.Get("reference")
	if versions.LessThan(cli.version, "1.25") && len(referenceFilters) > 0 {
		query.Set("filter", referenceFilters[0])
		for _, filterValue := range referenceFilters {
			optionFilters.Del("reference", filterValue)
		}
	}
	if optionFilters.Len() > 0 {
		//nolint:staticcheck // ignore SA1019 for old code
		filterJSON, err := filters.ToParamWithVersion(cli.version, optionFilters)
		if err != nil {
			return images, err
		}
		query.Set("filters", filterJSON)
	}
	if options.All {
		query.Set("all", "1")
	}
	if options.SharedSize && versions.GreaterThanOrEqualTo(cli.version, "1.42") {
		query.Set("shared-size", "1")
	}
	if options.Manifests && versions.GreaterThanOrEqualTo(cli.version, "1.47") {
		query.Set("manifests", "1")
	}

	resp, err := cli.get(ctx, "/images/json", query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return images, err
	}

	err = json.NewDecoder(resp.Body).Decode(&images)
	return images, err
}

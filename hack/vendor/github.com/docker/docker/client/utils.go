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
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	cerrdefs "github.com/containerd/errdefs"
	"github.com/docker/docker/api/types/filters"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

type emptyIDError string

func (e emptyIDError) InvalidParameter() {}

func (e emptyIDError) Error() string {
	return "invalid " + string(e) + " name or ID: value is empty"
}

// trimID trims the given object-ID / name, returning an error if it's empty.
func trimID(objType, id string) (string, error) {
	id = strings.TrimSpace(id)
	if id == "" {
		return "", emptyIDError(objType)
	}
	return id, nil
}

// getFiltersQuery returns a url query with "filters" query term, based on the
// filters provided.
func getFiltersQuery(f filters.Args) (url.Values, error) {
	query := url.Values{}
	if f.Len() > 0 {
		filterJSON, err := filters.ToJSON(f)
		if err != nil {
			return query, err
		}
		query.Set("filters", filterJSON)
	}
	return query, nil
}

// encodePlatforms marshals the given platform(s) to JSON format, to
// be used for query-parameters for filtering / selecting platforms.
func encodePlatforms(platform ...ocispec.Platform) ([]string, error) {
	if len(platform) == 0 {
		return []string{}, nil
	}
	if len(platform) == 1 {
		p, err := encodePlatform(&platform[0])
		if err != nil {
			return nil, err
		}
		return []string{p}, nil
	}

	seen := make(map[string]struct{}, len(platform))
	out := make([]string, 0, len(platform))
	for i := range platform {
		p, err := encodePlatform(&platform[i])
		if err != nil {
			return nil, err
		}
		if _, ok := seen[p]; !ok {
			out = append(out, p)
			seen[p] = struct{}{}
		}
	}
	return out, nil
}

// encodePlatform marshals the given platform to JSON format, to
// be used for query-parameters for filtering / selecting platforms. It
// is used as a helper for encodePlatforms,
func encodePlatform(platform *ocispec.Platform) (string, error) {
	p, err := json.Marshal(platform)
	if err != nil {
		return "", fmt.Errorf("%w: invalid platform: %v", cerrdefs.ErrInvalidArgument, err)
	}
	return string(p), nil
}

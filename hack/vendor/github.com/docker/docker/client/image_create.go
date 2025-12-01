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
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/distribution/reference"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
)

// ImageCreate creates a new image based on the parent options.
// It returns the JSON content in the response body.
func (cli *Client) ImageCreate(ctx context.Context, parentReference string, options image.CreateOptions) (io.ReadCloser, error) {
	ref, err := reference.ParseNormalizedNamed(parentReference)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Set("fromImage", ref.Name())
	query.Set("tag", getAPITagFromNamedRef(ref))
	if options.Platform != "" {
		query.Set("platform", strings.ToLower(options.Platform))
	}
	resp, err := cli.tryImageCreate(ctx, query, options.RegistryAuth)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (cli *Client) tryImageCreate(ctx context.Context, query url.Values, registryAuth string) (*http.Response, error) {
	return cli.post(ctx, "/images/create", query, nil, http.Header{
		registry.AuthHeader: {registryAuth},
	})
}

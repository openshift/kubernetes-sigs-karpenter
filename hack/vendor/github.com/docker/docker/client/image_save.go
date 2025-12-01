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
	"net/url"
)

// ImageSave retrieves one or more images from the docker host as an io.ReadCloser.
//
// Platforms is an optional parameter that specifies the platforms to save from the image.
// This is only has effect if the input image is a multi-platform image.
func (cli *Client) ImageSave(ctx context.Context, imageIDs []string, saveOpts ...ImageSaveOption) (io.ReadCloser, error) {
	var opts imageSaveOpts
	for _, opt := range saveOpts {
		if err := opt.Apply(&opts); err != nil {
			return nil, err
		}
	}

	query := url.Values{
		"names": imageIDs,
	}

	if len(opts.apiOptions.Platforms) > 0 {
		if err := cli.NewVersionError(ctx, "1.48", "platform"); err != nil {
			return nil, err
		}
		p, err := encodePlatforms(opts.apiOptions.Platforms...)
		if err != nil {
			return nil, err
		}
		query["platform"] = p
	}

	resp, err := cli.get(ctx, "/images/get", query, nil)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

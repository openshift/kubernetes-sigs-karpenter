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
	"net/url"

	"github.com/docker/docker/api/types/image"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// ImageHistoryWithPlatform sets the platform for the image history operation.
func ImageHistoryWithPlatform(platform ocispec.Platform) ImageHistoryOption {
	return imageHistoryOptionFunc(func(opt *imageHistoryOpts) error {
		if opt.apiOptions.Platform != nil {
			return fmt.Errorf("platform already set to %s", *opt.apiOptions.Platform)
		}
		opt.apiOptions.Platform = &platform
		return nil
	})
}

// ImageHistory returns the changes in an image in history format.
func (cli *Client) ImageHistory(ctx context.Context, imageID string, historyOpts ...ImageHistoryOption) ([]image.HistoryResponseItem, error) {
	query := url.Values{}

	var opts imageHistoryOpts
	for _, o := range historyOpts {
		if err := o.Apply(&opts); err != nil {
			return nil, err
		}
	}

	if opts.apiOptions.Platform != nil {
		if err := cli.NewVersionError(ctx, "1.48", "platform"); err != nil {
			return nil, err
		}

		p, err := encodePlatform(opts.apiOptions.Platform)
		if err != nil {
			return nil, err
		}
		query.Set("platform", p)
	}

	resp, err := cli.get(ctx, "/images/"+imageID+"/history", query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return nil, err
	}

	var history []image.HistoryResponseItem
	err = json.NewDecoder(resp.Body).Decode(&history)
	return history, err
}

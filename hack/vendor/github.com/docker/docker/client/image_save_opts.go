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
	"fmt"

	"github.com/docker/docker/api/types/image"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

type ImageSaveOption interface {
	Apply(*imageSaveOpts) error
}

type imageSaveOptionFunc func(opt *imageSaveOpts) error

func (f imageSaveOptionFunc) Apply(o *imageSaveOpts) error {
	return f(o)
}

// ImageSaveWithPlatforms sets the platforms to be saved from the image.
func ImageSaveWithPlatforms(platforms ...ocispec.Platform) ImageSaveOption {
	return imageSaveOptionFunc(func(opt *imageSaveOpts) error {
		if opt.apiOptions.Platforms != nil {
			return fmt.Errorf("platforms already set to %v", opt.apiOptions.Platforms)
		}
		opt.apiOptions.Platforms = platforms
		return nil
	})
}

type imageSaveOpts struct {
	apiOptions image.SaveOptions
}

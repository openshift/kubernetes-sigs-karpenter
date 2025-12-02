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

// ImageLoadOption is a type representing functional options for the image load operation.
type ImageLoadOption interface {
	Apply(*imageLoadOpts) error
}
type imageLoadOptionFunc func(opt *imageLoadOpts) error

func (f imageLoadOptionFunc) Apply(o *imageLoadOpts) error {
	return f(o)
}

type imageLoadOpts struct {
	apiOptions image.LoadOptions
}

// ImageLoadWithQuiet sets the quiet option for the image load operation.
func ImageLoadWithQuiet(quiet bool) ImageLoadOption {
	return imageLoadOptionFunc(func(opt *imageLoadOpts) error {
		opt.apiOptions.Quiet = quiet
		return nil
	})
}

// ImageLoadWithPlatforms sets the platforms to be loaded from the image.
func ImageLoadWithPlatforms(platforms ...ocispec.Platform) ImageLoadOption {
	return imageLoadOptionFunc(func(opt *imageLoadOpts) error {
		if opt.apiOptions.Platforms != nil {
			return fmt.Errorf("platforms already set to %v", opt.apiOptions.Platforms)
		}
		opt.apiOptions.Platforms = platforms
		return nil
	})
}

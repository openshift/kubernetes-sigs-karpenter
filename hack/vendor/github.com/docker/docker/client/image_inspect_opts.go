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
	"bytes"

	"github.com/docker/docker/api/types/image"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// ImageInspectOption is a type representing functional options for the image inspect operation.
type ImageInspectOption interface {
	Apply(*imageInspectOpts) error
}
type imageInspectOptionFunc func(opt *imageInspectOpts) error

func (f imageInspectOptionFunc) Apply(o *imageInspectOpts) error {
	return f(o)
}

// ImageInspectWithRawResponse instructs the client to additionally store the
// raw inspect response in the provided buffer.
func ImageInspectWithRawResponse(raw *bytes.Buffer) ImageInspectOption {
	return imageInspectOptionFunc(func(opts *imageInspectOpts) error {
		opts.raw = raw
		return nil
	})
}

// ImageInspectWithManifests sets manifests API option for the image inspect operation.
// This option is only available for API version 1.48 and up.
// With this option set, the image inspect operation response will have the
// [image.InspectResponse.Manifests] field populated if the server is multi-platform capable.
func ImageInspectWithManifests(manifests bool) ImageInspectOption {
	return imageInspectOptionFunc(func(clientOpts *imageInspectOpts) error {
		clientOpts.apiOptions.Manifests = manifests
		return nil
	})
}

// ImageInspectWithPlatform sets platform API option for the image inspect operation.
// This option is only available for API version 1.49 and up.
// With this option set, the image inspect operation will return information for the
// specified platform variant of the multi-platform image.
func ImageInspectWithPlatform(platform *ocispec.Platform) ImageInspectOption {
	return imageInspectOptionFunc(func(clientOpts *imageInspectOpts) error {
		clientOpts.apiOptions.Platform = platform
		return nil
	})
}

// ImageInspectWithAPIOpts sets the API options for the image inspect operation.
func ImageInspectWithAPIOpts(opts image.InspectOptions) ImageInspectOption {
	return imageInspectOptionFunc(func(clientOpts *imageInspectOpts) error {
		clientOpts.apiOptions = opts
		return nil
	})
}

type imageInspectOpts struct {
	raw        *bytes.Buffer
	apiOptions image.InspectOptions
}

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

import "net/http"

// NewClient initializes a new API client for the given host and API version.
// It uses the given http client as transport.
// It also initializes the custom http headers to add to each request.
//
// It won't send any version information if the version number is empty. It is
// highly recommended that you set a version or your client may break if the
// server is upgraded.
//
// Deprecated: use [NewClientWithOpts] passing the [WithHost], [WithVersion],
// [WithHTTPClient] and [WithHTTPHeaders] options. We recommend enabling API
// version negotiation by passing the [WithAPIVersionNegotiation] option instead
// of WithVersion.
func NewClient(host string, version string, client *http.Client, httpHeaders map[string]string) (*Client, error) {
	return NewClientWithOpts(WithHost(host), WithVersion(version), WithHTTPClient(client), WithHTTPHeaders(httpHeaders))
}

// NewEnvClient initializes a new API client based on environment variables.
// See FromEnv for a list of support environment variables.
//
// Deprecated: use [NewClientWithOpts] passing the [FromEnv] option.
func NewEnvClient() (*Client, error) {
	return NewClientWithOpts(FromEnv)
}

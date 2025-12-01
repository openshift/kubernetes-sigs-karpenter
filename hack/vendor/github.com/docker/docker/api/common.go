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

package api

// Common constants for daemon and client.
const (
	// DefaultVersion of the current REST API.
	DefaultVersion = "1.51"

	// MinSupportedAPIVersion is the minimum API version that can be supported
	// by the API server, specified as "major.minor". Note that the daemon
	// may be configured with a different minimum API version, as returned
	// in [github.com/docker/docker/api/types.Version.MinAPIVersion].
	//
	// API requests for API versions lower than the configured version produce
	// an error.
	MinSupportedAPIVersion = "1.24"

	// NoBaseImageSpecifier is the symbol used by the FROM
	// command to specify that no base image is to be used.
	NoBaseImageSpecifier = "scratch"
)

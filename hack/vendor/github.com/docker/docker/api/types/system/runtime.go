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

package system

// Runtime describes an OCI runtime
type Runtime struct {
	// "Legacy" runtime configuration for runc-compatible runtimes.

	Path string   `json:"path,omitempty"`
	Args []string `json:"runtimeArgs,omitempty"`

	// Shimv2 runtime configuration. Mutually exclusive with the legacy config above.

	Type    string                 `json:"runtimeType,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}

// RuntimeWithStatus extends [Runtime] to hold [RuntimeStatus].
type RuntimeWithStatus struct {
	Runtime
	Status map[string]string `json:"status,omitempty"`
}

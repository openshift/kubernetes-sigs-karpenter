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

package aws

// AccountIDEndpointMode controls how a resolved AWS account ID is handled for endpoint routing.
type AccountIDEndpointMode string

const (
	// AccountIDEndpointModeUnset indicates the AWS account ID will not be used for endpoint routing
	AccountIDEndpointModeUnset AccountIDEndpointMode = ""

	// AccountIDEndpointModePreferred indicates the AWS account ID will be used for endpoint routing if present
	AccountIDEndpointModePreferred = "preferred"

	// AccountIDEndpointModeRequired indicates an error will be returned if the AWS account ID is not resolved from identity
	AccountIDEndpointModeRequired = "required"

	// AccountIDEndpointModeDisabled indicates the AWS account ID will be ignored during endpoint routing
	AccountIDEndpointModeDisabled = "disabled"
)

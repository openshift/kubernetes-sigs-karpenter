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

package middleware

import (
	"github.com/aws/smithy-go/middleware"
)

// requestIDKey is used to retrieve request id from response metadata
type requestIDKey struct{}

// SetRequestIDMetadata sets the provided request id over middleware metadata
func SetRequestIDMetadata(metadata *middleware.Metadata, id string) {
	metadata.Set(requestIDKey{}, id)
}

// GetRequestIDMetadata retrieves the request id from middleware metadata
// returns string and bool indicating value of request id, whether request id was set.
func GetRequestIDMetadata(metadata middleware.Metadata) (string, bool) {
	if !metadata.Has(requestIDKey{}) {
		return "", false
	}

	v, ok := metadata.Get(requestIDKey{}).(string)
	if !ok {
		return "", true
	}
	return v, true
}

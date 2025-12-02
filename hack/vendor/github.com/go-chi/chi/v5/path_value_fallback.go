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

//go:build !go1.22 || tinygo
// +build !go1.22 tinygo

package chi

import "net/http"

// supportsPathValue is true if the Go version is 1.22 and above.
//
// If this is true, `net/http.Request` has methods `SetPathValue` and `PathValue`.
const supportsPathValue = false

// setPathValue sets the path values in the Request value
// based on the provided request context.
//
// setPathValue is only supported in Go 1.22 and above so
// this is just a blank function so that it compiles.
func setPathValue(rctx *Context, r *http.Request) {
}

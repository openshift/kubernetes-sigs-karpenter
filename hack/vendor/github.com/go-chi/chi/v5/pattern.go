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

//go:build go1.23 && !tinygo
// +build go1.23,!tinygo

package chi

import "net/http"

// supportsPattern is true if the Go version is 1.23 and above.
//
// If this is true, `net/http.Request` has field `Pattern`.
const supportsPattern = true

// setPattern sets the mux matched pattern in the http Request.
func setPattern(rctx *Context, r *http.Request) {
	r.Pattern = rctx.routePattern
}

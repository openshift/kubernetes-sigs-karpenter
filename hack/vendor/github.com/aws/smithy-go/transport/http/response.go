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

package http

import (
	"fmt"
	"net/http"
)

// Response provides the HTTP specific response structure for HTTP specific
// middleware steps to use to deserialize the response from an operation call.
type Response struct {
	*http.Response
}

// ResponseError provides the HTTP centric error type wrapping the underlying
// error with the HTTP response value.
type ResponseError struct {
	Response *Response
	Err      error
}

// HTTPStatusCode returns the HTTP response status code received from the service.
func (e *ResponseError) HTTPStatusCode() int { return e.Response.StatusCode }

// HTTPResponse returns the HTTP response received from the service.
func (e *ResponseError) HTTPResponse() *Response { return e.Response }

// Unwrap returns the nested error if any, or nil.
func (e *ResponseError) Unwrap() error { return e.Err }

func (e *ResponseError) Error() string {
	return fmt.Sprintf(
		"http response error StatusCode: %d, %v",
		e.Response.StatusCode, e.Err)
}

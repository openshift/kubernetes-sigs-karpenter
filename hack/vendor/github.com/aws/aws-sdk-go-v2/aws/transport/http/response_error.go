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
	"errors"
	"fmt"

	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// ResponseError provides the HTTP centric error type wrapping the underlying error
// with the HTTP response value and the deserialized RequestID.
type ResponseError struct {
	*smithyhttp.ResponseError

	// RequestID associated with response error
	RequestID string
}

// ServiceRequestID returns the request id associated with Response Error
func (e *ResponseError) ServiceRequestID() string { return e.RequestID }

// Error returns the formatted error
func (e *ResponseError) Error() string {
	return fmt.Sprintf(
		"https response error StatusCode: %d, RequestID: %s, %v",
		e.Response.StatusCode, e.RequestID, e.Err)
}

// As populates target and returns true if the type of target is a error type that
// the ResponseError embeds, (e.g.AWS HTTP ResponseError)
func (e *ResponseError) As(target interface{}) bool {
	return errors.As(e.ResponseError, target)
}

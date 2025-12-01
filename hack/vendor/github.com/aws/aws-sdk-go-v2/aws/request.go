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

import (
	"fmt"
)

// TODO remove replace with smithy.CanceledError

// RequestCanceledError is the error that will be returned by an API request
// that was canceled. Requests given a Context may return this error when
// canceled.
type RequestCanceledError struct {
	Err error
}

// CanceledError returns true to satisfy interfaces checking for canceled errors.
func (*RequestCanceledError) CanceledError() bool { return true }

// Unwrap returns the underlying error, if there was one.
func (e *RequestCanceledError) Unwrap() error {
	return e.Err
}
func (e *RequestCanceledError) Error() string {
	return fmt.Sprintf("request canceled, %v", e.Err)
}

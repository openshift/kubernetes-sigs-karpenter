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

package document

import (
	"fmt"
	"reflect"
)

// UnmarshalTypeError is an error type representing an error
// unmarshaling a Smithy document to a Go value type. This is different
// from UnmarshalError in that it does not wrap an underlying error type.
type UnmarshalTypeError struct {
	Value string
	Type  reflect.Type
}

// Error returns the string representation of the error.
// Satisfying the error interface.
func (e *UnmarshalTypeError) Error() string {
	return fmt.Sprintf("unmarshal failed, cannot unmarshal %s into Go value type %s",
		e.Value, e.Type.String())
}

// An InvalidUnmarshalError is an error type representing an invalid type
// encountered while unmarshaling a Smithy document to a Go value type.
type InvalidUnmarshalError struct {
	Type reflect.Type
}

// Error returns the string representation of the error.
// Satisfying the error interface.
func (e *InvalidUnmarshalError) Error() string {
	var msg string
	if e.Type == nil {
		msg = "cannot unmarshal to nil value"
	} else if e.Type.Kind() != reflect.Ptr {
		msg = fmt.Sprintf("cannot unmarshal to non-pointer value, got %s", e.Type.String())
	} else {
		msg = fmt.Sprintf("cannot unmarshal to nil value, %s", e.Type.String())
	}

	return fmt.Sprintf("unmarshal failed, %s", msg)
}

// An UnmarshalError wraps an error that occurred while unmarshaling a
// Smithy document into a Go type. This is different from
// UnmarshalTypeError in that it wraps the underlying error that occurred.
type UnmarshalError struct {
	Err   error
	Value string
	Type  reflect.Type
}

// Unwrap returns the underlying unmarshaling error
func (e *UnmarshalError) Unwrap() error {
	return e.Err
}

// Error returns the string representation of the error.
// Satisfying the error interface.
func (e *UnmarshalError) Error() string {
	return fmt.Sprintf("unmarshal failed, cannot unmarshal %q into %s, %v",
		e.Value, e.Type.String(), e.Err)
}

// An InvalidMarshalError is an error type representing an error
// occurring when marshaling a Go value type.
type InvalidMarshalError struct {
	Message string
}

// Error returns the string representation of the error.
// Satisfying the error interface.
func (e *InvalidMarshalError) Error() string {
	return fmt.Sprintf("marshal failed, %s", e.Message)
}

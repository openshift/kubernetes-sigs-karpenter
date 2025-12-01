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

package jsonpointer

type pointerError string

func (e pointerError) Error() string {
	return string(e)
}

const (
	// ErrPointer is an error raised by the jsonpointer package
	ErrPointer pointerError = "JSON pointer error"

	// ErrInvalidStart states that a JSON pointer must start with a separator ("/")
	ErrInvalidStart pointerError = `JSON pointer must be empty or start with a "` + pointerSeparator

	// ErrUnsupportedValueType indicates that a value of the wrong type is being set
	ErrUnsupportedValueType pointerError = "only structs, pointers, maps and slices are supported for setting values"
)

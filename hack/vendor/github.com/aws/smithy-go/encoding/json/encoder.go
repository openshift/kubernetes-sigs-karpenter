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

package json

import (
	"bytes"
)

// Encoder is JSON encoder that supports construction of JSON values
// using methods.
type Encoder struct {
	w *bytes.Buffer
	Value
}

// NewEncoder returns a new JSON encoder
func NewEncoder() *Encoder {
	writer := bytes.NewBuffer(nil)
	scratch := make([]byte, 64)

	return &Encoder{w: writer, Value: newValue(writer, &scratch)}
}

// String returns the String output of the JSON encoder
func (e Encoder) String() string {
	return e.w.String()
}

// Bytes returns the []byte slice of the JSON encoder
func (e Encoder) Bytes() []byte {
	return e.w.Bytes()
}

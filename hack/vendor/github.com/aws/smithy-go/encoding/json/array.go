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

// Array represents the encoding of a JSON Array
type Array struct {
	w          *bytes.Buffer
	writeComma bool
	scratch    *[]byte
}

func newArray(w *bytes.Buffer, scratch *[]byte) *Array {
	w.WriteRune(leftBracket)
	return &Array{w: w, scratch: scratch}
}

// Value adds a new element to the JSON Array.
// Returns a Value type that is used to encode
// the array element.
func (a *Array) Value() Value {
	if a.writeComma {
		a.w.WriteRune(comma)
	} else {
		a.writeComma = true
	}

	return newValue(a.w, a.scratch)
}

// Close encodes the end of the JSON Array
func (a *Array) Close() {
	a.w.WriteRune(rightBracket)
}

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

// Object represents the encoding of a JSON Object type
type Object struct {
	w          *bytes.Buffer
	writeComma bool
	scratch    *[]byte
}

func newObject(w *bytes.Buffer, scratch *[]byte) *Object {
	w.WriteRune(leftBrace)
	return &Object{w: w, scratch: scratch}
}

func (o *Object) writeKey(key string) {
	escapeStringBytes(o.w, []byte(key))
	o.w.WriteRune(colon)
}

// Key adds the given named key to the JSON object.
// Returns a Value encoder that should be used to encode
// a JSON value type.
func (o *Object) Key(name string) Value {
	if o.writeComma {
		o.w.WriteRune(comma)
	} else {
		o.writeComma = true
	}
	o.writeKey(name)
	return newValue(o.w, o.scratch)
}

// Close encodes the end of the JSON Object
func (o *Object) Close() {
	o.w.WriteRune(rightBrace)
}

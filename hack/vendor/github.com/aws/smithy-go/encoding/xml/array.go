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

package xml

// arrayMemberWrapper is the default member wrapper tag name for XML Array type
var arrayMemberWrapper = StartElement{
	Name: Name{Local: "member"},
}

// Array represents the encoding of a XML array type
type Array struct {
	w       writer
	scratch *[]byte

	// member start element is the array member wrapper start element
	memberStartElement StartElement

	// isFlattened indicates if the array is a flattened array.
	isFlattened bool
}

// newArray returns an array encoder.
// It also takes in the  member start element, array start element.
// It takes in a isFlattened bool, indicating that an array is flattened array.
//
// A wrapped array ["value1", "value2"] is represented as
// `<List><member>value1</member><member>value2</member></List>`.

// A flattened array `someList: ["value1", "value2"]` is represented as
// `<someList>value1</someList><someList>value2</someList>`.
func newArray(w writer, scratch *[]byte, memberStartElement StartElement, arrayStartElement StartElement, isFlattened bool) *Array {
	var memberWrapper = memberStartElement
	if isFlattened {
		memberWrapper = arrayStartElement
	}

	return &Array{
		w:                  w,
		scratch:            scratch,
		memberStartElement: memberWrapper,
		isFlattened:        isFlattened,
	}
}

// Member adds a new member to the XML array.
// It returns a Value encoder.
func (a *Array) Member() Value {
	v := newValue(a.w, a.scratch, a.memberStartElement)
	v.isFlattened = a.isFlattened
	return v
}

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

package flect

import (
	"strings"
)

// Pascalize returns a string with each segment capitalized
//	user = User
//	bob dylan = BobDylan
//	widget_id = WidgetID
func Pascalize(s string) string {
	return New(s).Pascalize().String()
}

// Pascalize returns a string with each segment capitalized
//	user = User
//	bob dylan = BobDylan
//	widget_id = WidgetID
func (i Ident) Pascalize() Ident {
	c := i.Camelize()
	if len(c.String()) == 0 {
		return c
	}
	if len(i.Parts) == 0 {
		return i
	}
	capLen := 1
	if _, ok := baseAcronyms[strings.ToUpper(i.Parts[0])]; ok {
		capLen = len(i.Parts[0])
	}
	return New(string(strings.ToUpper(c.Original[0:capLen])) + c.Original[capLen:])
}

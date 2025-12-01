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

import "unicode"

// Capitalize will cap the first letter of string
//	user = User
//	bob dylan = Bob dylan
//	widget_id = Widget_id
func Capitalize(s string) string {
	return New(s).Capitalize().String()
}

// Capitalize will cap the first letter of string
//	user = User
//	bob dylan = Bob dylan
//	widget_id = Widget_id
func (i Ident) Capitalize() Ident {
	if len(i.Parts) == 0 {
		return New("")
	}
	runes := []rune(i.Original)
	runes[0] = unicode.ToTitle(runes[0])
	return New(string(runes))
}

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
	"unicode"
)

// Dasherize returns an alphanumeric, lowercased, dashed string
//	Donald E. Knuth = donald-e-knuth
//	Test with + sign = test-with-sign
//	admin/WidgetID = admin-widget-id
func Dasherize(s string) string {
	return New(s).Dasherize().String()
}

// Dasherize returns an alphanumeric, lowercased, dashed string
//	Donald E. Knuth = donald-e-knuth
//	Test with + sign = test-with-sign
//	admin/WidgetID = admin-widget-id
func (i Ident) Dasherize() Ident {
	var parts []string

	for _, part := range i.Parts {
		var x string
		for _, c := range part {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				x += string(c)
			}
		}
		parts = xappend(parts, x)
	}

	return New(strings.ToLower(strings.Join(parts, "-")))
}

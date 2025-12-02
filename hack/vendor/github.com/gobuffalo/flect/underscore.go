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

// Underscore a string
//	bob dylan --> bob_dylan
//	Nice to see you! --> nice_to_see_you
//	widgetID --> widget_id
func Underscore(s string) string {
	return New(s).Underscore().String()
}

// Underscore a string
//	bob dylan --> bob_dylan
//	Nice to see you! --> nice_to_see_you
//	widgetID --> widget_id
func (i Ident) Underscore() Ident {
	out := make([]string, 0, len(i.Parts))
	for _, part := range i.Parts {
		var x strings.Builder
		x.Grow(len(part))
		for _, c := range part {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				x.WriteRune(c)
			}
		}
		if x.Len() > 0 {
			out = append(out, x.String())
		}
	}
	return New(strings.ToLower(strings.Join(out, "_")))
}

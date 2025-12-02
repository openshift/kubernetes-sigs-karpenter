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

// Camelize returns a camelize version of a string
//	bob dylan = bobDylan
//	widget_id = widgetID
//	WidgetID = widgetID
func Camelize(s string) string {
	return New(s).Camelize().String()
}

// Camelize returns a camelize version of a string
//	bob dylan = bobDylan
//	widget_id = widgetID
//	WidgetID = widgetID
func (i Ident) Camelize() Ident {
	var out []string
	for i, part := range i.Parts {
		var x string
		var capped bool
		for _, c := range part {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				if i == 0 {
					x += string(unicode.ToLower(c))
					continue
				}
				if !capped {
					capped = true
					x += string(unicode.ToUpper(c))
					continue
				}
				x += string(c)
			}
		}
		if x != "" {
			out = append(out, x)
		}
	}
	return New(strings.Join(out, ""))
}

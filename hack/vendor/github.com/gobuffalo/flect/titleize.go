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

// Titleize will capitalize the start of each part
//	"Nice to see you!" = "Nice To See You!"
//	"i've read a book! have you?" = "I've Read A Book! Have You?"
//	"This is `code` ok" = "This Is `code` OK"
func Titleize(s string) string {
	return New(s).Titleize().String()
}

// Titleize will capitalize the start of each part
//	"Nice to see you!" = "Nice To See You!"
//	"i've read a book! have you?" = "I've Read A Book! Have You?"
//	"This is `code` ok" = "This Is `code` OK"
func (i Ident) Titleize() Ident {
	var parts []string

	// TODO: we need to reconsider the design.
	//       this approach preserves inline code block as is but it also
	//       preserves the other words start with a special character.
	//       I would prefer: "*wonderful* world" to be "*Wonderful* World"
	for _, part := range i.Parts {
		// CAUTION: in unicode, []rune(str)[0] is not rune(str[0])
		runes := []rune(part)
		x := string(unicode.ToTitle(runes[0]))
		if len(runes) > 1 {
			x += string(runes[1:])
		}
		parts = append(parts, x)
	}

	return New(strings.Join(parts, " "))
}

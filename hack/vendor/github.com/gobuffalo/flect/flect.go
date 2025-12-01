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

/*
Package flect is a new inflection engine to replace [https://github.com/markbates/inflect](https://github.com/markbates/inflect) designed to be more modular, more readable, and easier to fix issues on than the original.
*/
package flect

import (
	"strings"
	"unicode"
)

var spaces = []rune{'_', ' ', ':', '-', '/'}

func isSpace(c rune) bool {
	for _, r := range spaces {
		if r == c {
			return true
		}
	}
	return unicode.IsSpace(c)
}

func xappend(a []string, ss ...string) []string {
	for _, s := range ss {
		s = strings.TrimSpace(s)
		for _, x := range spaces {
			s = strings.Trim(s, string(x))
		}
		if _, ok := baseAcronyms[strings.ToUpper(s)]; ok {
			s = strings.ToUpper(s)
		}
		if s != "" {
			a = append(a, s)
		}
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

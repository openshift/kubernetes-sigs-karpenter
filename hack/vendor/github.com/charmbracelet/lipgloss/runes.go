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

package lipgloss

import (
	"strings"
)

// StyleRunes apply a given style to runes at the given indices in the string.
// Note that you must provide styling options for both matched and unmatched
// runes. Indices out of bounds will be ignored.
func StyleRunes(str string, indices []int, matched, unmatched Style) string {
	// Convert slice of indices to a map for easier lookups
	m := make(map[int]struct{})
	for _, i := range indices {
		m[i] = struct{}{}
	}

	var (
		out   strings.Builder
		group strings.Builder
		style Style
		runes = []rune(str)
	)

	for i, r := range runes {
		group.WriteRune(r)

		_, matches := m[i]
		_, nextMatches := m[i+1]

		if matches != nextMatches || i == len(runes)-1 {
			// Flush
			if matches {
				style = matched
			} else {
				style = unmatched
			}
			out.WriteString(style.Render(group.String()))
			group.Reset()
		}
	}

	return out.String()
}

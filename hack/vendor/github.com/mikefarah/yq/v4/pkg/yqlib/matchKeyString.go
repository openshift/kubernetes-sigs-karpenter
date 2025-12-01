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

package yqlib

func matchKey(name string, pattern string) (matched bool) {
	if pattern == "" {
		return name == pattern
	}
	if pattern == "*" {
		log.Debug("wild!")
		return true
	}
	return deepMatch(name, pattern)
}

// deepMatch reports whether the name matches the pattern in linear time.
// Source https://research.swtch.com/glob
func deepMatch(name, pattern string) bool {
	px := 0
	nx := 0
	nextPx := 0
	nextNx := 0
	for px < len(pattern) || nx < len(name) {
		if px < len(pattern) {
			c := pattern[px]
			switch c {
			default: // ordinary character
				if nx < len(name) && name[nx] == c {
					px++
					nx++
					continue
				}
			case '?': // single-character wildcard
				if nx < len(name) {
					px++
					nx++
					continue
				}
			case '*': // zero-or-more-character wildcard
				// Try to match at nx.
				// If that doesn't work out,
				// restart at nx+1 next.
				nextPx = px
				nextNx = nx + 1
				px++
				continue
			}
		}
		// Mismatch. Maybe restart.
		if 0 < nextNx && nextNx <= len(name) {
			px = nextPx
			nx = nextNx
			continue
		}
		return false
	}
	// Matched all of pattern to all of name. Success.
	return true
}

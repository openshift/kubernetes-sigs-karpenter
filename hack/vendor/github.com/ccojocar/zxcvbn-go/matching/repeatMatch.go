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

package matching

import (
	"strings"

	"github.com/ccojocar/zxcvbn-go/entropy"
	"github.com/ccojocar/zxcvbn-go/match"
)

const repeatMatcherName = "REPEAT"

// FilterRepeatMatcher can be pass to zxcvbn-go.PasswordStrength to skip that matcher
func FilterRepeatMatcher(m match.Matcher) bool {
	return m.ID == repeatMatcherName
}

func repeatMatch(password string) []match.Match {
	var matches []match.Match

	// Loop through password. if current == prev currentStreak++ else if currentStreak > 2 {buildMatch; currentStreak = 1} prev = current
	var current, prev string
	currentStreak := 1
	var i int
	var char rune
	for i, char = range password {
		current = string(char)
		if i == 0 {
			prev = current
			continue
		}

		if strings.EqualFold(current, prev) {
			currentStreak++
		} else if currentStreak > 2 {
			iPos := i - currentStreak
			jPos := i - 1
			matchRepeat := match.Match{
				Pattern:        "repeat",
				I:              iPos,
				J:              jPos,
				Token:          password[iPos : jPos+1],
				DictionaryName: prev,
			}
			matchRepeat.Entropy = entropy.RepeatEntropy(matchRepeat)
			matches = append(matches, matchRepeat)
			currentStreak = 1
		} else {
			currentStreak = 1
		}

		prev = current
	}

	if currentStreak > 2 {
		iPos := i - currentStreak + 1
		jPos := i
		matchRepeat := match.Match{
			Pattern:        "repeat",
			I:              iPos,
			J:              jPos,
			Token:          password[iPos : jPos+1],
			DictionaryName: prev,
		}
		matchRepeat.Entropy = entropy.RepeatEntropy(matchRepeat)
		matches = append(matches, matchRepeat)
	}
	return matches
}

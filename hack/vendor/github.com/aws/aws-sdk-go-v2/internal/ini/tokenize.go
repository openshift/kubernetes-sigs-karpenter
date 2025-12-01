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

package ini

import (
	"strings"
)

func tokenize(lines []string) ([]lineToken, error) {
	tokens := make([]lineToken, 0, len(lines))
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 || isLineComment(line) {
			continue
		}

		if tok := asProfile(line); tok != nil {
			tokens = append(tokens, tok)
		} else if tok := asProperty(line); tok != nil {
			tokens = append(tokens, tok)
		} else if tok := asSubProperty(line); tok != nil {
			tokens = append(tokens, tok)
		} else if tok := asContinuation(line); tok != nil {
			tokens = append(tokens, tok)
		} // unrecognized tokens are effectively ignored
	}
	return tokens, nil
}

func isLineComment(line string) bool {
	trimmed := strings.TrimLeft(line, " \t")
	return strings.HasPrefix(trimmed, "#") || strings.HasPrefix(trimmed, ";")
}

func asProfile(line string) *lineTokenProfile { // " [ type name ] ; comment"
	trimmed := strings.TrimSpace(trimProfileComment(line)) // "[ type name ]"
	if !isBracketed(trimmed) {
		return nil
	}
	trimmed = trimmed[1 : len(trimmed)-1] // " type name " (or just " name ")
	trimmed = strings.TrimSpace(trimmed)  // "type name" / "name"
	typ, name := splitProfile(trimmed)
	return &lineTokenProfile{
		Type: typ,
		Name: name,
	}
}

func asProperty(line string) *lineTokenProperty {
	if isLineSpace(rune(line[0])) {
		return nil
	}

	trimmed := trimPropertyComment(line)
	trimmed = strings.TrimRight(trimmed, " \t")
	k, v, ok := splitProperty(trimmed)
	if !ok {
		return nil
	}

	return &lineTokenProperty{
		Key:   strings.ToLower(k), // LEGACY: normalize key case
		Value: legacyStrconv(v),   // LEGACY: see func docs
	}
}

func asSubProperty(line string) *lineTokenSubProperty {
	if !isLineSpace(rune(line[0])) {
		return nil
	}

	// comments on sub-properties are included in the value
	trimmed := strings.TrimLeft(line, " \t")
	k, v, ok := splitProperty(trimmed)
	if !ok {
		return nil
	}

	return &lineTokenSubProperty{ // same LEGACY constraints as in normal property
		Key:   strings.ToLower(k),
		Value: legacyStrconv(v),
	}
}

func asContinuation(line string) *lineTokenContinuation {
	if !isLineSpace(rune(line[0])) {
		return nil
	}

	// includes comments like sub-properties
	trimmed := strings.TrimLeft(line, " \t")
	return &lineTokenContinuation{
		Value: trimmed,
	}
}

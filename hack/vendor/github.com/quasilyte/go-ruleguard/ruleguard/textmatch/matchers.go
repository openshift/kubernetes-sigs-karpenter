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

package textmatch

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

// inputValue is a wrapper for string|[]byte.
//
// We hold both values to avoid string->[]byte and vice versa
// conversions when doing Match and MatchString.
type inputValue struct {
	s string
	b []byte
}

func newInputValue(s string) inputValue {
	return inputValue{s: s, b: []byte(s)}
}

type containsLiteralMatcher struct{ value inputValue }

func (m *containsLiteralMatcher) MatchString(s string) bool {
	return strings.Contains(s, m.value.s)
}

func (m *containsLiteralMatcher) Match(b []byte) bool {
	return bytes.Contains(b, m.value.b)
}

type prefixLiteralMatcher struct{ value inputValue }

func (m *prefixLiteralMatcher) MatchString(s string) bool {
	return strings.HasPrefix(s, m.value.s)
}

func (m *prefixLiteralMatcher) Match(b []byte) bool {
	return bytes.HasPrefix(b, m.value.b)
}

type suffixLiteralMatcher struct{ value inputValue }

func (m *suffixLiteralMatcher) MatchString(s string) bool {
	return strings.HasSuffix(s, m.value.s)
}

func (m *suffixLiteralMatcher) Match(b []byte) bool {
	return bytes.HasSuffix(b, m.value.b)
}

type eqLiteralMatcher struct{ value inputValue }

func (m *eqLiteralMatcher) MatchString(s string) bool {
	return m.value.s == s
}

func (m *eqLiteralMatcher) Match(b []byte) bool {
	return bytes.Equal(m.value.b, b)
}

type prefixRunePredMatcher struct{ pred func(rune) bool }

func (m *prefixRunePredMatcher) MatchString(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return m.pred(r)
}

func (m *prefixRunePredMatcher) Match(b []byte) bool {
	r, _ := utf8.DecodeRune(b)
	return m.pred(r)
}

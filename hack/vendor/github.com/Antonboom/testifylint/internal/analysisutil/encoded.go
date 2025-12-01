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

package analysisutil

import "strings"

var whitespaceRemover = strings.NewReplacer("\n", "", "\\n", "", "\t", "", "\\t", "", " ", "")

// IsJSONLike returns true if the string has JSON format features.
// A positive result can be returned for invalid JSON as well.
func IsJSONLike(s string) bool {
	s = whitespaceRemover.Replace(unescape(s))

	var startMatch bool
	for _, prefix := range []string{
		`{{`, `{[`, `{"`,
		`[{{`, `[{[`, `[{"`,
	} {
		if strings.HasPrefix(s, prefix) {
			startMatch = true
			break
		}
	}
	if !startMatch {
		return false
	}

	for _, keyValue := range []string{`":{`, `":[`, `":"`} {
		if strings.Contains(s, keyValue) {
			return true
		}
	}
	return false

	// NOTE(a.telyshev): We do not check the end of the string, because this is usually a field for typos.
	// And one of the reasons for using JSON-specific assertions is to catch typos like this.
}

func unescape(s string) string {
	s = strings.ReplaceAll(s, `\"`, `"`)
	s = unquote(s, `"`)
	s = unquote(s, "`")
	return s
}

func unquote(s string, q string) string {
	return strings.TrimLeft(strings.TrimRight(s, q), q)
}

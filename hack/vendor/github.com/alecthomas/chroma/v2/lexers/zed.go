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

package lexers

import (
	"strings"
)

// Zed lexer.
func init() { // nolint: gochecknoinits
	Get("Zed").SetAnalyser(func(text string) float32 {
		if strings.Contains(text, "definition ") && strings.Contains(text, "relation ") && strings.Contains(text, "permission ") {
			return 0.9
		}
		if strings.Contains(text, "definition ") {
			return 0.5
		}
		if strings.Contains(text, "relation ") {
			return 0.5
		}
		if strings.Contains(text, "permission ") {
			return 0.25
		}
		return 0.0
	})
}

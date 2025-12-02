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

// Package shared provides shared utilities for checkers.
package shared

import (
	"go/doc/comment"
	"strings"
)

// HasDeprecatedParagraph reports whether the given comment blocks contain a
// paragraph starting with deprecation marker.
func HasDeprecatedParagraph(blocks []comment.Block) bool {
	for _, block := range blocks {
		par, ok := block.(*comment.Paragraph)
		if !ok || len(par.Text) == 0 {
			continue
		}
		text, ok := (par.Text[0]).(comment.Plain)
		if !ok {
			continue
		}

		// Only an exact match (casing and the trailing whitespace) is considered
		// a valid deprecation marker.
		if strings.HasPrefix(string(text), "Deprecated: ") {
			return true
		}
	}
	return false
}

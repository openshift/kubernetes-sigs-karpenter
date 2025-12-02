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

package comment

import (
	"go/ast"
	"strings"
)

type Directive string

const (
	prefix                     = `//exhaustruct:`
	DirectiveIgnore  Directive = prefix + `ignore`
	DirectiveEnforce Directive = prefix + `enforce`
)

// HasDirective parses a directive from a given list of comments.
// If no directive is found, the second return value is `false`.
func HasDirective(comments []*ast.CommentGroup, expected Directive) bool {
	for _, cg := range comments {
		for _, commentLine := range cg.List {
			if strings.HasPrefix(commentLine.Text, string(expected)) {
				return true
			}
		}
	}

	return false
}

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

package analyzer

import (
	"go/ast"
	"strings"
)

const nolintPrefix = "//nolint"

func hasDisallowDirective(cg *ast.CommentGroup) bool {
	if cg == nil {
		return false
	}

	return directiveFound(cg)
}

func directiveFound(cg *ast.CommentGroup) bool {
	for i := len(cg.List) - 1; i >= 0; i-- {
		comment := cg.List[i]
		if !strings.HasPrefix(comment.Text, nolintPrefix) {
			continue
		}

		startingIdx := len(nolintPrefix)
		for {
			idx := strings.Index(comment.Text[startingIdx:], name)
			if idx == -1 {
				break
			}

			if len(comment.Text[startingIdx+idx:]) == len(name) {
				return true
			}

			c := comment.Text[startingIdx+idx+len(name)]
			if c == '.' || c == ',' || c == ' ' || c == '	' {
				return true
			}
			startingIdx += idx + 1
		}
	}

	return false
}

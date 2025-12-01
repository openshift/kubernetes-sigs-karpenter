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

package exhaustive

import (
	"go/ast"
	"go/token"
	"strings"
)

const (
	ignoreComment                     = "//exhaustive:ignore"
	enforceComment                    = "//exhaustive:enforce"
	ignoreDefaultCaseRequiredComment  = "//exhaustive:ignore-default-case-required"
	enforceDefaultCaseRequiredComment = "//exhaustive:enforce-default-case-required"
)

func hasCommentPrefix(comments []*ast.CommentGroup, comment string) bool {
	for _, c := range comments {
		for _, cc := range c.List {
			if strings.HasPrefix(cc.Text, comment) {
				return true
			}
		}
	}
	return false
}

func fileCommentMap(fset *token.FileSet, file *ast.File) ast.CommentMap {
	return ast.NewCommentMap(fset, file, file.Comments)
}

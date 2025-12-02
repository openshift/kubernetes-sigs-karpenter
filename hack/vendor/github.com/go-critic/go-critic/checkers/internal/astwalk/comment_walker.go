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

package astwalk

import (
	"go/ast"
	"strings"
)

type commentWalker struct {
	visitor CommentVisitor
}

func (w *commentWalker) WalkFile(f *ast.File) {
	if !w.visitor.EnterFile(f) {
		return
	}

	for _, cg := range f.Comments {
		visitCommentGroups(cg, w.visitor.VisitComment)
	}
}

func visitCommentGroups(cg *ast.CommentGroup, visit func(*ast.CommentGroup)) {
	var group []*ast.Comment
	visitGroup := func(list []*ast.Comment) {
		if len(list) == 0 {
			return
		}
		cg := &ast.CommentGroup{List: list}
		visit(cg)
	}
	for _, comment := range cg.List {
		if strings.HasPrefix(comment.Text, "/*") {
			visitGroup(group)
			group = group[:0]
			visitGroup([]*ast.Comment{comment})
		} else {
			group = append(group, comment)
		}
	}
	visitGroup(group)
}

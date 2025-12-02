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
)

type localCommentWalker struct {
	visitor LocalCommentVisitor
}

func (w *localCommentWalker) WalkFile(f *ast.File) {
	if !w.visitor.EnterFile(f) {
		return
	}

	for _, decl := range f.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok || !w.visitor.EnterFunc(decl) {
			continue
		}

		for _, cg := range f.Comments {
			// Not sure that decls/comments are sorted
			// by positions, so do a naive full scan for now.
			if cg.Pos() < decl.Pos() || cg.Pos() > decl.End() {
				continue
			}

			visitCommentGroups(cg, w.visitor.VisitLocalComment)
		}
	}
}

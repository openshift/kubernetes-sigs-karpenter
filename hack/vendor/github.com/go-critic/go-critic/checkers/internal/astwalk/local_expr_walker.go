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

type localExprWalker struct {
	visitor LocalExprVisitor
}

func (w *localExprWalker) WalkFile(f *ast.File) {
	if !w.visitor.EnterFile(f) {
		return
	}

	for _, decl := range f.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok || !w.visitor.EnterFunc(decl) {
			continue
		}
		ast.Inspect(decl.Body, func(x ast.Node) bool {
			if x, ok := x.(ast.Expr); ok {
				w.visitor.VisitLocalExpr(x)
				return !w.visitor.skipChilds()
			}
			return true
		})
	}
}

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

type docCommentWalker struct {
	visitor DocCommentVisitor
}

func (w *docCommentWalker) WalkFile(f *ast.File) {
	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			if decl.Doc != nil {
				w.visitor.VisitDocComment(decl.Doc)
			}
		case *ast.GenDecl:
			if decl.Doc != nil {
				w.visitor.VisitDocComment(decl.Doc)
			}
			for _, spec := range decl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					if spec.Doc != nil {
						w.visitor.VisitDocComment(spec.Doc)
					}
				case *ast.ValueSpec:
					if spec.Doc != nil {
						w.visitor.VisitDocComment(spec.Doc)
					}
				case *ast.TypeSpec:
					if spec.Doc != nil {
						w.visitor.VisitDocComment(spec.Doc)
					}
					ast.Inspect(spec.Type, func(n ast.Node) bool {
						if n, ok := n.(*ast.Field); ok {
							if n.Doc != nil {
								w.visitor.VisitDocComment(n.Doc)
							}
						}
						return true
					})
				}
			}
		}
	}
}

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

package dogsled

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.DogsledSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name: "dogsled",
			Doc:  "Checks assignments with too many blank identifiers (e.g. x, _, _, _, := f())",
			Run: func(pass *analysis.Pass) (any, error) {
				return run(pass, settings.MaxBlankIdentifiers)
			},
			Requires: []*analysis.Analyzer{inspect.Analyzer},
		}).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

func run(pass *analysis.Pass, maxBlanks int) (any, error) {
	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	for node := range insp.PreorderSeq((*ast.FuncDecl)(nil)) {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if funcDecl.Body == nil {
			continue
		}

		for _, expr := range funcDecl.Body.List {
			assgnStmt, ok := expr.(*ast.AssignStmt)
			if !ok {
				continue
			}

			numBlank := 0
			for _, left := range assgnStmt.Lhs {
				ident, ok := left.(*ast.Ident)
				if !ok {
					continue
				}
				if ident.Name == "_" {
					numBlank++
				}
			}

			if numBlank > maxBlanks {
				pass.Reportf(assgnStmt.Pos(), "declaration has %v blank identifiers", numBlank)
			}
		}
	}

	return nil, nil
}

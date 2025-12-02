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

package gochecknoinits

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name:     "gochecknoinits",
			Doc:      "Checks that no init functions are present in Go code",
			Run:      run,
			Requires: []*analysis.Analyzer{inspect.Analyzer},
		}).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

func run(pass *analysis.Pass) (any, error) {
	insp, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, nil
	}

	for node := range insp.PreorderSeq((*ast.FuncDecl)(nil)) {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			continue
		}

		fnName := funcDecl.Name.Name
		if fnName == "init" && funcDecl.Recv.NumFields() == 0 {
			pass.Reportf(funcDecl.Pos(), "don't use %s function", internal.FormatCode(fnName))
		}
	}

	return nil, nil
}

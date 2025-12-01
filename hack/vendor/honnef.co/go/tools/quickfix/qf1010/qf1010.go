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

package qf1010

import (
	"go/ast"
	"go/types"

	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/edit"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/analysis/report"
	"honnef.co/go/tools/knowledge"
	"honnef.co/go/tools/pattern"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "QF1010",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title:    "Convert slice of bytes to string when printing it",
		Since:    "2021.1",
		Severity: lint.SeverityHint,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var byteSlicePrintingQ = pattern.MustParse(`
	(Or
		(CallExpr
			(Symbol (Or
				"fmt.Print"
				"fmt.Println"
				"fmt.Sprint"
				"fmt.Sprintln"
				"log.Fatal"
				"log.Fatalln"
				"log.Panic"
				"log.Panicln"
				"log.Print"
				"log.Println"
				"(*log.Logger).Fatal"
				"(*log.Logger).Fatalln"
				"(*log.Logger).Panic"
				"(*log.Logger).Panicln"
				"(*log.Logger).Print"
				"(*log.Logger).Println")) args)

		(CallExpr (Symbol (Or
			"fmt.Fprint"
			"fmt.Fprintln")) _:args))`)

var byteSlicePrintingR = pattern.MustParse(`(CallExpr (Ident "string") [arg])`)

func run(pass *analysis.Pass) (interface{}, error) {
	fn := func(node ast.Node) {
		m, ok := code.Match(pass, byteSlicePrintingQ, node)
		if !ok {
			return
		}
		args := m.State["args"].([]ast.Expr)
		for _, arg := range args {
			if !code.IsOfStringConvertibleByteSlice(pass, arg) {
				continue
			}
			if types.Implements(pass.TypesInfo.TypeOf(arg), knowledge.Interfaces["fmt.Stringer"]) {
				continue
			}

			fix := edit.Fix("Convert argument to string", edit.ReplaceWithPattern(pass.Fset, arg, byteSlicePrintingR, pattern.State{"arg": arg}))
			report.Report(pass, arg, "could convert argument to string", report.Fixes(fix))
		}
	}
	code.Preorder(pass, fn, (*ast.CallExpr)(nil))
	return nil, nil
}

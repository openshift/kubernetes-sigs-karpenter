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

package s1033

import (
	"go/ast"

	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/edit"
	"honnef.co/go/tools/analysis/facts/generated"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/analysis/report"
	"honnef.co/go/tools/pattern"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "S1033",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer, generated.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title:   `Unnecessary guard around call to \"delete\"`,
		Text:    `Calling \'delete\' on a nil map is a no-op.`,
		Since:   "2019.2",
		MergeIf: lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var checkGuardedDeleteQ = pattern.MustParse(`
	(IfStmt
		(AssignStmt
			[(Ident "_") ok@(Ident _)]
			":="
			(IndexExpr m key))
		ok
		[call@(CallExpr (Builtin "delete") [m key])]
		nil)`)

func run(pass *analysis.Pass) (interface{}, error) {
	fn := func(node ast.Node) {
		if m, ok := code.Match(pass, checkGuardedDeleteQ, node); ok {
			report.Report(pass, node, "unnecessary guard around call to delete",
				report.ShortRange(),
				report.FilterGenerated(),
				report.Fixes(edit.Fix("remove guard", edit.ReplaceWithNode(pass.Fset, node, m.State["call"].(ast.Node)))))
		}
	}

	code.Preorder(pass, fn, (*ast.IfStmt)(nil))
	return nil, nil
}

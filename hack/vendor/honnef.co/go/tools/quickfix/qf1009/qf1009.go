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

package qf1009

import (
	"go/ast"
	"go/token"

	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/edit"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/analysis/report"
	"honnef.co/go/tools/pattern"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "QF1009",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title:    `Use \'time.Time.Equal\' instead of \'==\' operator`,
		Since:    "2021.1",
		Severity: lint.SeverityInfo,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var timeEqualR = pattern.MustParse(`(CallExpr (SelectorExpr lhs (Ident "Equal")) rhs)`)

func run(pass *analysis.Pass) (interface{}, error) {
	// FIXME(dh): create proper suggested fix for renamed import

	fn := func(node ast.Node) {
		expr := node.(*ast.BinaryExpr)
		if expr.Op != token.EQL {
			return
		}
		if !code.IsOfTypeWithName(pass, expr.X, "time.Time") || !code.IsOfTypeWithName(pass, expr.Y, "time.Time") {
			return
		}
		report.Report(pass, node, "probably want to use time.Time.Equal instead",
			report.Fixes(edit.Fix("Use time.Time.Equal method",
				edit.ReplaceWithPattern(pass.Fset, node, timeEqualR, pattern.State{"lhs": expr.X, "rhs": expr.Y}))))
	}
	code.Preorder(pass, fn, (*ast.BinaryExpr)(nil))
	return nil, nil
}

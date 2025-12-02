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

package sa4013

import (
	"go/ast"

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
		Name:     "SA4013",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title:    `Negating a boolean twice (\'!!b\') is the same as writing \'b\'. This is either redundant, or a typo.`,
		Since:    "2017.1",
		Severity: lint.SeverityWarning,
		MergeIf:  lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var checkDoubleNegationQ = pattern.MustParse(`(UnaryExpr "!" single@(UnaryExpr "!" x))`)

func run(pass *analysis.Pass) (interface{}, error) {
	fn := func(node ast.Node) {
		if m, ok := code.Match(pass, checkDoubleNegationQ, node); ok {
			report.Report(pass, node, "negating a boolean twice has no effect; is this a typo?", report.Fixes(
				edit.Fix("turn into single negation", edit.ReplaceWithNode(pass.Fset, node, m.State["single"].(ast.Node))),
				edit.Fix("remove double negation", edit.ReplaceWithNode(pass.Fset, node, m.State["x"].(ast.Node)))))
		}
	}
	code.Preorder(pass, fn, (*ast.UnaryExpr)(nil))
	return nil, nil
}

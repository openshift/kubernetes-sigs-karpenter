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

package sa1013

import (
	"go/ast"

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
		Name:     "SA1013",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title:    `\'io.Seeker.Seek\' is being called with the whence constant as the first argument, but it should be the second`,
		Since:    "2017.1",
		Severity: lint.SeverityWarning,
		MergeIf:  lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var (
	checkSeekerQ = pattern.MustParse(`(CallExpr fun@(SelectorExpr _ (Ident "Seek")) [arg1@(SelectorExpr _ (Symbol (Or "io.SeekStart" "io.SeekCurrent" "io.SeekEnd"))) arg2])`)
	checkSeekerR = pattern.MustParse(`(CallExpr fun [arg2 arg1])`)
)

func run(pass *analysis.Pass) (interface{}, error) {
	fn := func(node ast.Node) {
		if m, edits, ok := code.MatchAndEdit(pass, checkSeekerQ, checkSeekerR, node); ok {
			if !code.IsMethod(pass, m.State["fun"].(*ast.SelectorExpr), "Seek", knowledge.Signatures["(io.Seeker).Seek"]) {
				return
			}
			report.Report(pass, node, "the first argument of io.Seeker is the offset, but an io.Seek* constant is being used instead",
				report.Fixes(edit.Fix("swap arguments", edits...)))
		}
	}
	code.Preorder(pass, fn, (*ast.CallExpr)(nil))
	return nil, nil
}

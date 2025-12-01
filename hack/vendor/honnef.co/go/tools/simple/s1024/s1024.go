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

package s1024

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
		Name:     "S1024",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer, generated.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title: `Replace \'x.Sub(time.Now())\' with \'time.Until(x)\'`,
		Text: `The \'time.Until\' helper has the same effect as using \'x.Sub(time.Now())\'
but is easier to read.`,
		Before:  `x.Sub(time.Now())`,
		After:   `time.Until(x)`,
		Since:   "2017.1",
		MergeIf: lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var (
	checkTimeUntilQ = pattern.MustParse(`(CallExpr (Symbol "(time.Time).Sub") [(CallExpr (Symbol "time.Now") [])])`)
	checkTimeUntilR = pattern.MustParse(`(CallExpr (SelectorExpr (Ident "time") (Ident "Until")) [arg])`)
)

func run(pass *analysis.Pass) (interface{}, error) {
	fn := func(node ast.Node) {
		if _, ok := code.Match(pass, checkTimeUntilQ, node); ok {
			if sel, ok := node.(*ast.CallExpr).Fun.(*ast.SelectorExpr); ok {
				r := pattern.NodeToAST(checkTimeUntilR.Root, map[string]interface{}{"arg": sel.X}).(ast.Node)
				report.Report(pass, node, "should use time.Until instead of t.Sub(time.Now())",
					report.FilterGenerated(),
					report.MinimumStdlibVersion("go1.8"),
					report.Fixes(edit.Fix("replace with call to time.Until", edit.ReplaceWithNode(pass.Fset, node, r))))
			} else {
				report.Report(pass, node, "should use time.Until instead of t.Sub(time.Now())",
					report.MinimumStdlibVersion("go1.8"),
					report.FilterGenerated())
			}
		}
	}
	code.Preorder(pass, fn, (*ast.CallExpr)(nil))
	return nil, nil
}

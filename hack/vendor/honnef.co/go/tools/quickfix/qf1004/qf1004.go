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

package qf1004

import (
	"fmt"
	"go/ast"
	"go/types"

	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/edit"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/analysis/report"
	"honnef.co/go/tools/go/types/typeutil"
	"honnef.co/go/tools/pattern"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "QF1004",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title:    `Use \'strings.ReplaceAll\' instead of \'strings.Replace\' with \'n == -1\'`,
		Since:    "2021.1",
		Severity: lint.SeverityHint,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var stringsReplaceAllQ = pattern.MustParse(`(Or
	(CallExpr fn@(Symbol "strings.Replace") [_ _ _ lit@(IntegerLiteral "-1")])
	(CallExpr fn@(Symbol "strings.SplitN") [_ _ lit@(IntegerLiteral "-1")])
	(CallExpr fn@(Symbol "strings.SplitAfterN") [_ _ lit@(IntegerLiteral "-1")])
	(CallExpr fn@(Symbol "bytes.Replace") [_ _ _ lit@(IntegerLiteral "-1")])
	(CallExpr fn@(Symbol "bytes.SplitN") [_ _ lit@(IntegerLiteral "-1")])
	(CallExpr fn@(Symbol "bytes.SplitAfterN") [_ _ lit@(IntegerLiteral "-1")]))`)

func run(pass *analysis.Pass) (interface{}, error) {
	// XXX respect minimum Go version

	// FIXME(dh): create proper suggested fix for renamed import

	fn := func(node ast.Node) {
		matcher, ok := code.Match(pass, stringsReplaceAllQ, node)
		if !ok {
			return
		}

		var replacement string
		switch typeutil.FuncName(matcher.State["fn"].(*types.Func)) {
		case "strings.Replace":
			replacement = "strings.ReplaceAll"
		case "strings.SplitN":
			replacement = "strings.Split"
		case "strings.SplitAfterN":
			replacement = "strings.SplitAfter"
		case "bytes.Replace":
			replacement = "bytes.ReplaceAll"
		case "bytes.SplitN":
			replacement = "bytes.Split"
		case "bytes.SplitAfterN":
			replacement = "bytes.SplitAfter"
		default:
			panic("unreachable")
		}

		call := node.(*ast.CallExpr)
		report.Report(pass, call.Fun, fmt.Sprintf("could use %s instead", replacement),
			report.Fixes(edit.Fix(fmt.Sprintf("Use %s instead", replacement),
				edit.ReplaceWithString(call.Fun, replacement),
				edit.Delete(matcher.State["lit"].(ast.Node)))))
	}
	code.Preorder(pass, fn, (*ast.CallExpr)(nil))
	return nil, nil
}

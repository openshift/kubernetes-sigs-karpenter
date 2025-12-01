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

package sa4024

import (
	"fmt"
	"go/ast"

	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/analysis/report"
	"honnef.co/go/tools/pattern"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA4024",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title: `Checking for impossible return value from a builtin function`,
		Text: `Return values of the \'len\' and \'cap\' builtins cannot be negative.

See https://golang.org/pkg/builtin/#len and https://golang.org/pkg/builtin/#cap.

Example:

    if len(slice) < 0 {
        fmt.Println("unreachable code")
    }`,
		Since:    "2021.1",
		Severity: lint.SeverityWarning,
		MergeIf:  lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var builtinLessThanZeroQ = pattern.MustParse(`
	(Or
		(BinaryExpr
			(IntegerLiteral "0")
			">"
			(CallExpr builtin@(Builtin (Or "len" "cap")) _))
		(BinaryExpr
			(CallExpr builtin@(Builtin (Or "len" "cap")) _)
			"<"
			(IntegerLiteral "0")))
`)

func run(pass *analysis.Pass) (interface{}, error) {
	fn := func(node ast.Node) {
		matcher, ok := code.Match(pass, builtinLessThanZeroQ, node)
		if !ok {
			return
		}

		builtin := matcher.State["builtin"].(*ast.Ident)
		report.Report(pass, node, fmt.Sprintf("builtin function %s does not return negative values", builtin.Name))
	}
	code.Preorder(pass, fn, (*ast.BinaryExpr)(nil))

	return nil, nil
}

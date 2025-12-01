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

package sa3001

import (
	"fmt"
	"go/ast"

	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/analysis/report"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA3001",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title: `Assigning to \'b.N\' in benchmarks distorts the results`,
		Text: `The testing package dynamically sets \'b.N\' to improve the reliability of
benchmarks and uses it in computations to determine the duration of a
single operation. Benchmark code must not alter \'b.N\' as this would
falsify results.`,
		Since:    "2017.1",
		Severity: lint.SeverityError,
		MergeIf:  lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

func run(pass *analysis.Pass) (interface{}, error) {
	fn := func(node ast.Node) {
		assign := node.(*ast.AssignStmt)
		if len(assign.Lhs) != 1 || len(assign.Rhs) != 1 {
			return
		}
		sel, ok := assign.Lhs[0].(*ast.SelectorExpr)
		if !ok {
			return
		}
		if sel.Sel.Name != "N" {
			return
		}
		if !code.IsOfPointerToTypeWithName(pass, sel.X, "testing.B") {
			return
		}
		report.Report(pass, assign, fmt.Sprintf("should not assign to %s", report.Render(pass, sel)))
	}
	code.Preorder(pass, fn, (*ast.AssignStmt)(nil))
	return nil, nil
}

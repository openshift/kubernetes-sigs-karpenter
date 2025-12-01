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

package sa4028

import (
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
		Name:     "SA4028",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title:    `\'x % 1\' is always zero`,
		Since:    "2022.1",
		Severity: lint.SeverityWarning,
		MergeIf:  lint.MergeIfAny, // MergeIfAny if we only flag literals, not named constants
	},
})

var Analyzer = SCAnalyzer.Analyzer

var moduloOneQ = pattern.MustParse(`(BinaryExpr _ "%" (IntegerLiteral "1"))`)

func run(pass *analysis.Pass) (interface{}, error) {
	fn := func(node ast.Node) {
		_, ok := code.Match(pass, moduloOneQ, node)
		if !ok {
			return
		}
		report.Report(pass, node, "x % 1 is always zero")
	}
	code.Preorder(pass, fn, (*ast.BinaryExpr)(nil))
	return nil, nil
}

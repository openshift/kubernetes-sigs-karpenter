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

package sa4001

import (
	"go/ast"
	"regexp"

	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/analysis/report"
	"honnef.co/go/tools/pattern"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA4001",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title:    `\'&*x\' gets simplified to \'x\', it does not copy \'x\'`,
		Since:    "2017.1",
		Severity: lint.SeverityWarning,
		MergeIf:  lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var (
	// cgo produces code like fn(&*_Cvar_kSomeCallbacks) which we don't
	// want to flag.
	cgoIdent               = regexp.MustCompile(`^_C(func|var)_.+$`)
	checkIneffectiveCopyQ1 = pattern.MustParse(`(UnaryExpr "&" (StarExpr obj))`)
	checkIneffectiveCopyQ2 = pattern.MustParse(`(StarExpr (UnaryExpr "&" _))`)
)

func run(pass *analysis.Pass) (interface{}, error) {
	fn := func(node ast.Node) {
		if m, ok := code.Match(pass, checkIneffectiveCopyQ1, node); ok {
			if ident, ok := m.State["obj"].(*ast.Ident); !ok || !cgoIdent.MatchString(ident.Name) {
				report.Report(pass, node, "&*x will be simplified to x. It will not copy x.")
			}
		} else if _, ok := code.Match(pass, checkIneffectiveCopyQ2, node); ok {
			report.Report(pass, node, "*&x will be simplified to x. It will not copy x.")
		}
	}
	code.Preorder(pass, fn, (*ast.UnaryExpr)(nil), (*ast.StarExpr)(nil))
	return nil, nil
}

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

package st1015

import (
	"go/ast"
	"go/token"

	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/facts/generated"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/analysis/report"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "ST1015",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer, generated.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title:   `A switch's default case should be the first or last case`,
		Since:   "2019.1",
		MergeIf: lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

func run(pass *analysis.Pass) (interface{}, error) {
	hasFallthrough := func(clause ast.Stmt) bool {
		// A valid fallthrough statement may be used only as the final non-empty statement in a case clause. Thus we can
		// easily avoid falsely matching fallthroughs in nested switches by not descending into blocks.

		body := clause.(*ast.CaseClause).Body
		for i := len(body) - 1; i >= 0; i-- {
			last := body[i]
			switch stmt := last.(type) {
			case *ast.EmptyStmt:
				// Fallthrough may be followed by empty statements
			case *ast.BranchStmt:
				return stmt.Tok == token.FALLTHROUGH
			default:
				return false
			}
		}

		return false
	}

	fn := func(node ast.Node) {
		stmt := node.(*ast.SwitchStmt)
		list := stmt.Body.List
		defaultIdx := -1
		for i, c := range list {
			if c.(*ast.CaseClause).List == nil {
				defaultIdx = i
				break
			}
		}

		if defaultIdx == -1 || defaultIdx == 0 || defaultIdx == len(list)-1 {
			// No default case, or it's the first or last case
			return
		}

		if hasFallthrough(list[defaultIdx-1]) || hasFallthrough(list[defaultIdx]) {
			// We either fall into or out of this case; don't mess with the order
			return
		}

		report.Report(pass, list[defaultIdx], "default case should be first or last in switch statement", report.FilterGenerated())
	}
	code.Preorder(pass, fn, (*ast.SwitchStmt)(nil))
	return nil, nil
}

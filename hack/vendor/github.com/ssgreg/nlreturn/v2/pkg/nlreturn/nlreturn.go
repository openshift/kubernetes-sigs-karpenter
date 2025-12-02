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

package nlreturn

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

const (
	linterName = "nlreturn"
	linterDoc  = `Linter requires a new line before return and branch statements except when the return is alone inside a statement group (such as an if statement) to increase code clarity.`
)

var blockSize int

// NewAnalyzer returns a new nlreturn analyzer.
func NewAnalyzer() *analysis.Analyzer {
	a := &analysis.Analyzer{
		Name: linterName,
		Doc:  linterDoc,
		Run:  run,
	}

	a.Flags.Init("nlreturn", flag.ExitOnError)
	a.Flags.IntVar(&blockSize, "block-size", 1, "set block size that is still ok")

	return a
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			switch c := node.(type) {
			case *ast.CaseClause:
				inspectBlock(pass, c.Body)
			case *ast.CommClause:
				inspectBlock(pass, c.Body)
			case *ast.BlockStmt:
				inspectBlock(pass, c.List)
			}

			return true
		})
	}

	return nil, nil
}

func inspectBlock(pass *analysis.Pass, block []ast.Stmt) {
	for i, stmt := range block {
		switch stmt.(type) {
		case *ast.BranchStmt, *ast.ReturnStmt:

			if i == 0 || line(pass, stmt.Pos())-line(pass, block[0].Pos()) < blockSize {
				return
			}

			if line(pass, stmt.Pos())-line(pass, block[i-1].End()) <= 1 {
				pass.Report(analysis.Diagnostic{
					Pos:     stmt.Pos(),
					Message: fmt.Sprintf("%s with no blank line before", name(stmt)),
					SuggestedFixes: []analysis.SuggestedFix{
						{
							TextEdits: []analysis.TextEdit{
								{
									Pos:     stmt.Pos(),
									NewText: []byte("\n"),
									End:     stmt.Pos(),
								},
							},
						},
					},
				})
			}
		}
	}
}

func name(stmt ast.Stmt) string {
	switch c := stmt.(type) {
	case *ast.BranchStmt:
		return c.Tok.String()
	case *ast.ReturnStmt:
		return "return"
	default:
		return "unknown"
	}
}

func line(pass *analysis.Pass, pos token.Pos) int {
	return pass.Fset.Position(pos).Line
}

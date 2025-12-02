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

package checks

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"

	config "github.com/tommy-muehle/go-mnd/v2/config"
)

const CaseCheck = "case"

type CaseAnalyzer struct {
	pass   *analysis.Pass
	config *config.Config
}

func NewCaseAnalyzer(pass *analysis.Pass, config *config.Config) *CaseAnalyzer {
	return &CaseAnalyzer{
		pass:   pass,
		config: config,
	}
}

func (a *CaseAnalyzer) NodeFilter() []ast.Node {
	return []ast.Node{
		(*ast.CaseClause)(nil),
	}
}

func (a *CaseAnalyzer) Check(n ast.Node) {
	caseClause, ok := n.(*ast.CaseClause)
	if !ok {
		return
	}

	for _, c := range caseClause.List {
		switch x := c.(type) {
		case *ast.BasicLit:
			if a.isMagicNumber(x) {
				a.pass.Reportf(x.Pos(), reportMsg, x.Value, CaseCheck)
			}
		case *ast.BinaryExpr:
			a.checkBinaryExpr(x)
		}
	}
}

func (a *CaseAnalyzer) checkBinaryExpr(expr *ast.BinaryExpr) {
	switch x := expr.X.(type) {
	case *ast.BasicLit:
		if a.isMagicNumber(x) {
			a.pass.Reportf(x.Pos(), reportMsg, x.Value, CaseCheck)
		}
	}

	switch y := expr.Y.(type) {
	case *ast.BasicLit:
		if a.isMagicNumber(y) {
			a.pass.Reportf(y.Pos(), reportMsg, y.Value, CaseCheck)
		}
	}
}

func (a *CaseAnalyzer) isMagicNumber(l *ast.BasicLit) bool {
	return (l.Kind == token.FLOAT || l.Kind == token.INT) && !a.config.IsIgnoredNumber(l.Value)
}

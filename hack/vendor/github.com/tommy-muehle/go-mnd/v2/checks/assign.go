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

const AssignCheck = "assign"

type AssignAnalyzer struct {
	pass   *analysis.Pass
	config *config.Config
}

func NewAssignAnalyzer(pass *analysis.Pass, config *config.Config) *AssignAnalyzer {
	return &AssignAnalyzer{
		pass:   pass,
		config: config,
	}
}

func (a *AssignAnalyzer) NodeFilter() []ast.Node {
	return []ast.Node{
		(*ast.KeyValueExpr)(nil),
		(*ast.AssignStmt)(nil),
	}
}

func (a *AssignAnalyzer) Check(n ast.Node) {
	switch expr := n.(type) {
	case *ast.KeyValueExpr:
		switch x := expr.Value.(type) {
		case *ast.BasicLit:
			if a.isMagicNumber(x) {
				a.pass.Reportf(x.Pos(), reportMsg, x.Value, AssignCheck)
			}
		case *ast.BinaryExpr:
			a.checkBinaryExpr(x)
		}
	case *ast.AssignStmt:
		for _, e := range expr.Rhs {
			switch y := e.(type) {
			case *ast.UnaryExpr:
				a.checkUnaryExpr(y)
			case *ast.BinaryExpr:
				switch x := y.Y.(type) {
				case *ast.UnaryExpr:
					a.checkUnaryExpr(x)
				}
			}
		}
	}
}

func (a *AssignAnalyzer) checkUnaryExpr(expr *ast.UnaryExpr) {
	switch x := expr.X.(type) {
	case *ast.BasicLit:
		if a.isMagicNumber(x) {
			a.pass.Reportf(x.Pos(), reportMsg, x.Value, AssignCheck)
		}
	}
}

func (a *AssignAnalyzer) checkBinaryExpr(expr *ast.BinaryExpr) {
	switch x := expr.X.(type) {
	case *ast.BasicLit:
		if a.isMagicNumber(x) {
			a.pass.Reportf(x.Pos(), reportMsg, x.Value, AssignCheck)
		}
	}

	switch y := expr.Y.(type) {
	case *ast.BasicLit:
		if a.isMagicNumber(y) {
			a.pass.Reportf(y.Pos(), reportMsg, y.Value, AssignCheck)
		}
	}
}

func (a *AssignAnalyzer) isMagicNumber(l *ast.BasicLit) bool {
	return (l.Kind == token.FLOAT || l.Kind == token.INT) && !a.config.IsIgnoredNumber(l.Value)
}

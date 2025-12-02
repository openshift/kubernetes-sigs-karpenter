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

package rule

import (
	"go/ast"
	"go/token"

	"github.com/mgechev/revive/lint"
)

// BoolLiteralRule warns when logic expressions contains Boolean literals.
type BoolLiteralRule struct{}

// Apply applies the rule to given file.
func (*BoolLiteralRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}

	astFile := file.AST
	w := &lintBoolLiteral{astFile, onFailure}
	ast.Walk(w, astFile)

	return failures
}

// Name returns the rule name.
func (*BoolLiteralRule) Name() string {
	return "bool-literal-in-expr"
}

type lintBoolLiteral struct {
	file      *ast.File
	onFailure func(lint.Failure)
}

func (w *lintBoolLiteral) Visit(node ast.Node) ast.Visitor {
	if n, ok := node.(*ast.BinaryExpr); ok {
		if !isBoolOp(n.Op) {
			return w
		}

		lexeme, ok := isExprABooleanLit(n.X)
		if !ok {
			lexeme, ok = isExprABooleanLit(n.Y)
			if !ok {
				return w
			}
		}

		isConstant := (n.Op == token.LAND && lexeme == "false") || (n.Op == token.LOR && lexeme == "true")

		if isConstant {
			w.addFailure(n, "Boolean expression seems to always evaluate to "+lexeme, lint.FailureCategoryLogic)
		} else {
			w.addFailure(n, "omit Boolean literal in expression", lint.FailureCategoryStyle)
		}
	}

	return w
}

func (w lintBoolLiteral) addFailure(node ast.Node, msg string, cat lint.FailureCategory) {
	w.onFailure(lint.Failure{
		Confidence: 1,
		Node:       node,
		Category:   cat,
		Failure:    msg,
	})
}

// isBoolOp returns true if the given token corresponds to a bool operator.
func isBoolOp(t token.Token) bool {
	switch t {
	case token.LAND, token.LOR, token.EQL, token.NEQ:
		return true
	}

	return false
}

func isExprABooleanLit(n ast.Node) (lexeme string, ok bool) {
	oper, ok := n.(*ast.Ident)

	if !ok {
		return "", false
	}

	return oper.Name, oper.Name == "true" || oper.Name == "false"
}

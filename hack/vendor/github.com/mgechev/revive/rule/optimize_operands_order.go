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
	"fmt"
	"go/ast"
	"go/token"

	"github.com/mgechev/revive/internal/astutils"
	"github.com/mgechev/revive/lint"
)

// OptimizeOperandsOrderRule checks inefficient conditional expressions.
type OptimizeOperandsOrderRule struct{}

// Apply applies the rule to given file.
func (*OptimizeOperandsOrderRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}
	w := lintOptimizeOperandsOrderExpr{
		onFailure: onFailure,
	}
	ast.Walk(w, file.AST)
	return failures
}

// Name returns the rule name.
func (*OptimizeOperandsOrderRule) Name() string {
	return "optimize-operands-order"
}

type lintOptimizeOperandsOrderExpr struct {
	onFailure func(failure lint.Failure)
}

// Visit checks boolean AND and OR expressions to determine
// if swapping their operands may result in an execution speedup.
func (w lintOptimizeOperandsOrderExpr) Visit(node ast.Node) ast.Visitor {
	binExpr, ok := node.(*ast.BinaryExpr)
	if !ok {
		return w
	}

	switch binExpr.Op {
	case token.LAND, token.LOR:
	default:
		return w
	}

	isCaller := func(n ast.Node) bool {
		ce, ok := n.(*ast.CallExpr)
		if !ok {
			return false
		}

		ident, isIdent := ce.Fun.(*ast.Ident)
		if !isIdent {
			return true
		}

		return ident.Name != "len" || ident.Obj != nil
	}

	// check if the left sub-expression contains a function call
	call := astutils.SeekNode[*ast.CallExpr](binExpr.X, isCaller)
	if call == nil {
		return w
	}

	// check if the right sub-expression does not contain a function call
	call = astutils.SeekNode[*ast.CallExpr](binExpr.Y, isCaller)
	if call != nil {
		return w
	}

	newExpr := ast.BinaryExpr{X: binExpr.Y, Y: binExpr.X, Op: binExpr.Op}
	w.onFailure(lint.Failure{
		Failure:    fmt.Sprintf("for better performance '%v' might be rewritten as '%v'", astutils.GoFmt(binExpr), astutils.GoFmt(&newExpr)),
		Node:       node,
		Category:   lint.FailureCategoryOptimization,
		Confidence: 0.3,
	})

	return w
}

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

package checkers

import (
	"go/ast"
	"go/token"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"

	"github.com/go-toolsmith/astequal"
	"github.com/go-toolsmith/typep"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "nilValReturn"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects return statements those results evaluate to nil"
	info.Before = `
if err == nil {
	return err
}`
	info.After = `
// (A) - return nil explicitly
if err == nil {
	return nil
}
// (B) - typo in "==", change to "!="
if err != nil {
	return err
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&nilValReturnChecker{ctx: ctx}), nil
	})
}

type nilValReturnChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *nilValReturnChecker) VisitStmt(stmt ast.Stmt) {
	ifStmt, ok := stmt.(*ast.IfStmt)
	if !ok || len(ifStmt.Body.List) != 1 {
		return
	}
	ret, ok := ifStmt.Body.List[0].(*ast.ReturnStmt)
	if !ok {
		return
	}
	expr, ok := ifStmt.Cond.(*ast.BinaryExpr)
	if !ok {
		return
	}
	xIsNil := expr.Op == token.EQL &&
		typep.SideEffectFree(c.ctx.TypesInfo, expr.X) &&
		qualifiedName(expr.Y) == "nil"
	if !xIsNil {
		return
	}
	for _, res := range ret.Results {
		if astequal.Expr(expr.X, res) {
			c.warn(ret, expr.X)
			break
		}
	}
}

func (c *nilValReturnChecker) warn(cause, val ast.Node) {
	c.ctx.Warn(cause, "returned expr is always nil; replace %s with nil", val)
}

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

	"github.com/go-toolsmith/astcast"
	"github.com/go-toolsmith/astcopy"
	"github.com/go-toolsmith/typep"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "methodExprCall"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects method expression call that can be replaced with a method call"
	info.Before = `f := foo{}
foo.bar(f)`
	info.After = `f := foo{}
f.bar()`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&methodExprCallChecker{ctx: ctx}), nil
	})
}

type methodExprCallChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *methodExprCallChecker) VisitExpr(x ast.Expr) {
	call := astcast.ToCallExpr(x)
	s := astcast.ToSelectorExpr(call.Fun)

	if len(call.Args) < 1 || astcast.ToIdent(call.Args[0]).Name == "nil" {
		return
	}

	if typep.IsTypeExpr(c.ctx.TypesInfo, s.X) {
		c.warn(call, s)
	}
}

func (c *methodExprCallChecker) warn(cause *ast.CallExpr, s *ast.SelectorExpr) {
	selector := astcopy.SelectorExpr(s)
	selector.X = cause.Args[0]

	// Remove "&" from the receiver (if any).
	if u, ok := selector.X.(*ast.UnaryExpr); ok && u.Op == token.AND {
		selector.X = u.X
	}

	c.ctx.Warn(cause, "consider to change `%s` to `%s`", cause.Fun, selector)
}

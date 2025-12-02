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
	"go/types"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/checkers/internal/lintutil"
	"github.com/go-critic/go-critic/linter"

	"github.com/go-toolsmith/astcast"
	"golang.org/x/tools/go/ast/astutil"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "newDeref"
	info.Tags = []string{linter.StyleTag}
	info.Summary = "Detects immediate dereferencing of `new` expressions"
	info.Before = `x := *new(bool)`
	info.After = `x := false`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&newDerefChecker{ctx: ctx}), nil
	})
}

type newDerefChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *newDerefChecker) VisitExpr(expr ast.Expr) {
	deref := astcast.ToStarExpr(expr)
	call := astcast.ToCallExpr(deref.X)
	if astcast.ToIdent(call.Fun).Name == "new" {
		typ := c.ctx.TypeOf(call.Args[0])
		// allow *new(T) if T is a type parameter, see #1272 for details
		if _, ok := typ.(*types.TypeParam); ok {
			return
		}
		zv := lintutil.ZeroValueOf(astutil.Unparen(call.Args[0]), typ)
		if zv != nil {
			c.warn(expr, zv)
		}
	}
}

func (c *newDerefChecker) warn(cause, suggestion ast.Expr) {
	c.ctx.Warn(cause, "replace `%s` with `%s`", cause, suggestion)
}

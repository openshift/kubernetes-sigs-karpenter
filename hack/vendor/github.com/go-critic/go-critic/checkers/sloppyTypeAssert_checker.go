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
	"github.com/go-critic/go-critic/linter"

	"github.com/go-toolsmith/astcast"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "sloppyTypeAssert"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects redundant type assertions"
	info.Before = `
func f(r io.Reader) interface{} {
	return r.(interface{})
}
`
	info.After = `
func f(r io.Reader) interface{} {
	return r
}
`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&sloppyTypeAssertChecker{ctx: ctx}), nil
	})
}

type sloppyTypeAssertChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *sloppyTypeAssertChecker) VisitExpr(expr ast.Expr) {
	assert := astcast.ToTypeAssertExpr(expr)
	if assert.Type == nil {
		return
	}

	toType := c.ctx.TypeOf(expr)
	fromType := c.ctx.TypeOf(assert.X)

	if types.Identical(toType, fromType) {
		c.warnIdentical(expr)
		return
	}
}

func (c *sloppyTypeAssertChecker) warnIdentical(cause ast.Expr) {
	c.ctx.Warn(cause, "type assertion from/to types are identical")
}

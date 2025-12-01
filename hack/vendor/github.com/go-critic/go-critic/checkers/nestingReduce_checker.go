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

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "nestingReduce"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag, linter.ExperimentalTag}
	info.Params = linter.CheckerParams{
		"bodyWidth": {
			Value: 5,
			Usage: "min number of statements inside a branch to trigger a warning",
		},
	}
	info.Summary = "Finds where nesting level could be reduced"
	info.Before = `
for _, v := range a {
	if v.Bool {
		body()
	}
}`
	info.After = `
for _, v := range a {
	if !v.Bool {
		continue
	}
	body()
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &nestingReduceChecker{ctx: ctx}
		c.bodyWidth = info.Params.Int("bodyWidth")
		return astwalk.WalkerForStmt(c), nil
	})
}

type nestingReduceChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	bodyWidth int
}

func (c *nestingReduceChecker) VisitStmt(stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.ForStmt:
		c.checkLoopBody(stmt.Body.List)
	case *ast.RangeStmt:
		c.checkLoopBody(stmt.Body.List)
	}
}

func (c *nestingReduceChecker) checkLoopBody(body []ast.Stmt) {
	if len(body) != 1 {
		return
	}
	stmt, ok := body[0].(*ast.IfStmt)
	if !ok {
		return
	}
	if len(stmt.Body.List) >= c.bodyWidth && stmt.Else == nil {
		c.warnLoop(stmt)
	}
}

func (c *nestingReduceChecker) warnLoop(cause ast.Node) {
	c.ctx.Warn(cause, "invert if cond, replace body with `continue`, move old body after the statement")
}

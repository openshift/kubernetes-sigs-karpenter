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

	"github.com/go-toolsmith/astp"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "initClause"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag, linter.ExperimentalTag}
	info.Summary = "Detects non-assignment statements inside if/switch init clause"
	info.Before = `if sideEffect(); cond {
}`
	info.After = `sideEffect()
if cond {
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&initClauseChecker{ctx: ctx}), nil
	})
}

type initClauseChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *initClauseChecker) VisitStmt(stmt ast.Stmt) {
	initClause := c.getInitClause(stmt)
	if initClause != nil && !astp.IsAssignStmt(initClause) {
		c.warn(stmt, initClause)
	}
}

func (c *initClauseChecker) getInitClause(x ast.Stmt) ast.Stmt {
	switch x := x.(type) {
	case *ast.IfStmt:
		return x.Init
	case *ast.SwitchStmt:
		return x.Init
	default:
		return nil
	}
}

func (c *initClauseChecker) warn(stmt, clause ast.Stmt) {
	name := "if"
	if astp.IsSwitchStmt(stmt) {
		name = "switch"
	}
	c.ctx.Warn(stmt, "consider to move `%s` before %s", clause, name)
}

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

	"github.com/go-toolsmith/astequal"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "dupBranchBody"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects duplicated branch bodies inside conditional statements"
	info.Before = `
if cond {
	println("cond=true")
} else {
	println("cond=true")
}`
	info.After = `
if cond {
	println("cond=true")
} else {
	println("cond=false")
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&dupBranchBodyChecker{ctx: ctx}), nil
	})
}

type dupBranchBodyChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *dupBranchBodyChecker) VisitStmt(stmt ast.Stmt) {
	// TODO(quasilyte): extend to check switch statements as well.
	// Should be very careful with type switches.

	if stmt, ok := stmt.(*ast.IfStmt); ok {
		c.checkIf(stmt)
	}
}

func (c *dupBranchBodyChecker) checkIf(stmt *ast.IfStmt) {
	thenBody := stmt.Body
	elseBody, ok := stmt.Else.(*ast.BlockStmt)
	if ok && astequal.Stmt(thenBody, elseBody) {
		c.warnIf(stmt)
	}
}

func (c *dupBranchBodyChecker) warnIf(cause ast.Node) {
	c.ctx.Warn(cause, "both branches in if statement have same body")
}

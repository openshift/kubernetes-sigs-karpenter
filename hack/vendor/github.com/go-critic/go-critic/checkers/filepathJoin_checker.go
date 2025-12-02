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
	"strings"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"

	"github.com/go-toolsmith/astcast"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "filepathJoin"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects problems in filepath.Join() function calls"
	info.Before = `filepath.Join("dir/", filename)`
	info.After = `filepath.Join("dir", filename)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&filepathJoinChecker{ctx: ctx}), nil
	})
}

type filepathJoinChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *filepathJoinChecker) VisitExpr(expr ast.Expr) {
	call := astcast.ToCallExpr(expr)
	if qualifiedName(call.Fun) != "filepath.Join" {
		return
	}

	for _, arg := range call.Args {
		arg, ok := arg.(*ast.BasicLit)
		if ok && c.hasSeparator(arg) {
			c.warnSeparator(arg)
		}
	}
}

func (c *filepathJoinChecker) hasSeparator(v *ast.BasicLit) bool {
	return strings.ContainsAny(v.Value, `/\`)
}

func (c *filepathJoinChecker) warnSeparator(sep ast.Expr) {
	c.ctx.Warn(sep, "%s contains a path separator", sep)
}

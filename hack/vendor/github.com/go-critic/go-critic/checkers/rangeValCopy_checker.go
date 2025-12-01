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
	info.Name = "rangeValCopy"
	info.Tags = []string{linter.PerformanceTag}
	info.Params = linter.CheckerParams{
		"sizeThreshold": {
			Value: 128,
			Usage: "size in bytes that makes the warning trigger",
		},
		"skipTestFuncs": {
			Value: true,
			Usage: "whether to check test functions",
		},
	}
	info.Summary = "Detects loops that copy big objects during each iteration"
	info.Details = "Suggests to use index access or take address and make use pointer instead."
	info.Before = `
xs := make([][1024]byte, length)
for _, x := range xs {
	// Loop body.
}`
	info.After = `
xs := make([][1024]byte, length)
for i := range xs {
	x := &xs[i]
	// Loop body.
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &rangeValCopyChecker{ctx: ctx}
		c.sizeThreshold = int64(info.Params.Int("sizeThreshold"))
		c.skipTestFuncs = info.Params.Bool("skipTestFuncs")
		return astwalk.WalkerForStmt(c), nil
	})
}

type rangeValCopyChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	sizeThreshold int64
	skipTestFuncs bool
}

func (c *rangeValCopyChecker) EnterFunc(fn *ast.FuncDecl) bool {
	return fn.Body != nil &&
		!(c.skipTestFuncs && isUnitTestFunc(c.ctx, fn))
}

func (c *rangeValCopyChecker) VisitStmt(stmt ast.Stmt) {
	rng, ok := stmt.(*ast.RangeStmt)
	if !ok || rng.Value == nil {
		return
	}
	typ := c.ctx.TypeOf(rng.Value)
	if typ == nil {
		return
	}
	size, ok := c.ctx.SizeOf(typ)
	if ok && size >= c.sizeThreshold {
		c.warn(rng, size)
	}
}

func (c *rangeValCopyChecker) warn(n ast.Node, size int64) {
	c.ctx.Warn(n, "each iteration copies %d bytes (consider pointers or indexing)", size)
}

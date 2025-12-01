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
	info.Name = "captLocal"
	info.Tags = []string{linter.StyleTag}
	info.Params = linter.CheckerParams{
		"paramsOnly": {
			Value: true,
			Usage: "whether to restrict checker to params only",
		},
	}
	info.Summary = "Detects capitalized names for local variables"
	info.Before = `func f(IN int, OUT *int) (ERR error) {}`
	info.After = `func f(in int, out *int) (err error) {}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &captLocalChecker{ctx: ctx}
		c.paramsOnly = info.Params.Bool("paramsOnly")
		return astwalk.WalkerForLocalDef(c, ctx.TypesInfo), nil
	})
}

type captLocalChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	paramsOnly bool
}

func (c *captLocalChecker) VisitLocalDef(def astwalk.Name, _ ast.Expr) {
	if c.paramsOnly && def.Kind != astwalk.NameParam {
		return
	}
	if ast.IsExported(def.ID.Name) {
		c.warn(def.ID)
	}
}

func (c *captLocalChecker) warn(id ast.Node) {
	c.ctx.Warn(id, "`%s' should not be capitalized", id)
}

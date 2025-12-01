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
	info.Name = "builtinShadow"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag}
	info.Summary = "Detects when predeclared identifiers are shadowed in assignments"
	info.Before = `len := 10`
	info.After = `length := 10`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForLocalDef(&builtinShadowChecker{ctx: ctx}, ctx.TypesInfo), nil
	})
}

type builtinShadowChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *builtinShadowChecker) VisitLocalDef(name astwalk.Name, _ ast.Expr) {
	if isBuiltin(name.ID.Name) {
		c.warn(name.ID)
	}
}

func (c *builtinShadowChecker) warn(ident *ast.Ident) {
	c.ctx.Warn(ident, "shadowing of predeclared identifier: %s", ident)
}

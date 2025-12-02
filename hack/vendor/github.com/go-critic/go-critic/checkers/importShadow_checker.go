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
)

func init() {
	var info linter.CheckerInfo
	info.Name = "importShadow"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag}
	info.Summary = "Detects when imported package names shadowed in the assignments"
	info.Before = `
// "path/filepath" is imported.
filepath := "foo.txt"`
	info.After = `
filename := "foo.txt"`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		ctx.Require.PkgObjects = true
		return astwalk.WalkerForLocalDef(&importShadowChecker{ctx: ctx}, ctx.TypesInfo), nil
	})
}

type importShadowChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *importShadowChecker) VisitLocalDef(def astwalk.Name, _ ast.Expr) {
	for pkgObj, name := range c.ctx.PkgObjects {
		if name == def.ID.Name && name != "_" {
			c.warn(def.ID, name, pkgObj.Imported())
		}
	}
}

func (c *importShadowChecker) warn(id ast.Node, importedName string, pkg *types.Package) {
	if isStdlibPkg(pkg) {
		c.ctx.Warn(id, "shadow of imported package '%s'", importedName)
	} else {
		c.ctx.Warn(id, "shadow of imported from '%s' package '%s'", pkg.Path(), importedName)
	}
}

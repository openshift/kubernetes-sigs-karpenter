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
	info.Name = "tooManyResultsChecker"
	info.Tags = []string{linter.StyleTag, linter.OpinionatedTag, linter.ExperimentalTag}
	info.Params = linter.CheckerParams{
		"maxResults": {
			Value: 5,
			Usage: "maximum number of results",
		},
	}
	info.Summary = "Detects function with too many results"
	info.Before = `func fn() (a, b, c, d float32, _ int, _ bool)`
	info.After = `func fn() (resultStruct, bool)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := astwalk.WalkerForFuncDecl(&tooManyResultsChecker{
			ctx:       ctx,
			maxParams: info.Params.Int("maxResults"),
		})
		return c, nil
	})
}

type tooManyResultsChecker struct {
	astwalk.WalkHandler
	ctx       *linter.CheckerContext
	maxParams int
}

func (c *tooManyResultsChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	typ := c.ctx.TypeOf(decl.Name)
	sig, ok := typ.(*types.Signature)
	if !ok {
		return
	}

	if count := sig.Results().Len(); count > c.maxParams {
		c.warn(decl)
	}
}

func (c *tooManyResultsChecker) warn(n ast.Node) {
	c.ctx.Warn(n, "function has more than %d results, consider to simplify the function", c.maxParams)
}

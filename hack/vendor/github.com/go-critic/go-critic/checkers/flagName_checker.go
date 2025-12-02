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
	"go/constant"
	"go/types"
	"strings"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"

	"github.com/go-toolsmith/astcast"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "flagName"
	info.Tags = []string{linter.DiagnosticTag}
	info.Summary = "Detects suspicious flag names"
	info.Before = `b := flag.Bool(" foo ", false, "description")`
	info.After = `b := flag.Bool("foo", false, "description")`
	info.Note = "https://github.com/golang/go/issues/41792"

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&flagNameChecker{ctx: ctx}), nil
	})
}

type flagNameChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *flagNameChecker) VisitExpr(expr ast.Expr) {
	call := astcast.ToCallExpr(expr)
	calledExpr := astcast.ToSelectorExpr(call.Fun)
	obj, ok := c.ctx.TypesInfo.ObjectOf(astcast.ToIdent(calledExpr.X)).(*types.PkgName)
	if !ok {
		return
	}
	sym := calledExpr.Sel
	pkg := obj.Imported()
	if pkg.Path() != "flag" {
		return
	}

	switch sym.Name {
	case "Bool", "Duration", "Float64", "String",
		"Int", "Int64", "Uint", "Uint64":
		c.checkFlagName(call, call.Args[0])
	case "BoolVar", "DurationVar", "Float64Var", "StringVar",
		"IntVar", "Int64Var", "UintVar", "Uint64Var":
		c.checkFlagName(call, call.Args[1])
	}
}

func (c *flagNameChecker) checkFlagName(call *ast.CallExpr, arg ast.Expr) {
	cv := c.ctx.TypesInfo.Types[arg].Value
	if cv == nil {
		return // Non-constant name
	}
	name := constant.StringVal(cv)
	switch {
	case name == "":
		c.warnEmpty(call)
	case strings.HasPrefix(name, "-"):
		c.warnHyphenPrefix(call, name)
	case strings.Contains(name, "="):
		c.warnEq(call, name)
	case strings.Contains(name, " "):
		c.warnWhitespace(call, name)
	}
}

func (c *flagNameChecker) warnEmpty(cause ast.Node) {
	c.ctx.Warn(cause, "empty flag name")
}

func (c *flagNameChecker) warnHyphenPrefix(cause ast.Node, name string) {
	c.ctx.Warn(cause, "flag name %q should not start with a hyphen", name)
}

func (c *flagNameChecker) warnEq(cause ast.Node, name string) {
	c.ctx.Warn(cause, "flag name %q should not contain '='", name)
}

func (c *flagNameChecker) warnWhitespace(cause ast.Node, name string) {
	c.ctx.Warn(cause, "flag name %q contains whitespace", name)
}

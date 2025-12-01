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

package actual

import (
	"go/ast"
	gotypes "go/types"

	"golang.org/x/tools/go/analysis"

	"github.com/nunnatsa/ginkgolinter/internal/gomegahandler"
	"github.com/nunnatsa/ginkgolinter/internal/gomegainfo"
)

type Actual struct {
	Orig         *ast.CallExpr
	Clone        *ast.CallExpr
	Arg          ArgPayload
	argType      gotypes.Type
	isTuple      bool
	isAsync      bool
	asyncArg     *AsyncArg
	actualOffset int
}

func New(origExpr, cloneExpr *ast.CallExpr, orig *ast.CallExpr, clone *ast.CallExpr, pass *analysis.Pass, timePkg string, info *gomegahandler.GomegaBasicInfo) (*Actual, bool) {
	arg, actualOffset := getActualArgPayload(orig, clone, pass, info)
	if arg == nil {
		return nil, false
	}

	argType := pass.TypesInfo.TypeOf(orig.Args[actualOffset])
	isTuple := false

	if tpl, ok := argType.(*gotypes.Tuple); ok {
		if tpl.Len() > 0 {
			argType = tpl.At(0).Type()
		} else {
			argType = nil
		}

		isTuple = tpl.Len() > 1
	}

	isAsyncExpr := gomegainfo.IsAsyncActualMethod(info.MethodName)

	var asyncArg *AsyncArg
	if isAsyncExpr {
		asyncArg = newAsyncArg(origExpr, cloneExpr, orig, clone, argType, pass, actualOffset, timePkg)
	}

	return &Actual{
		Orig:         orig,
		Clone:        clone,
		Arg:          arg,
		argType:      argType,
		isTuple:      isTuple,
		isAsync:      isAsyncExpr,
		asyncArg:     asyncArg,
		actualOffset: actualOffset,
	}, true
}

func (a *Actual) ReplaceActual(newArgs ast.Expr) {
	a.Clone.Args[a.actualOffset] = newArgs
}

func (a *Actual) ReplaceActualWithItsFirstArg() {
	firstArgOfArg := a.Clone.Args[a.actualOffset].(*ast.CallExpr).Args[0]
	a.ReplaceActual(firstArgOfArg)
}

func (a *Actual) IsAsync() bool {
	return a.isAsync
}

func (a *Actual) IsTuple() bool {
	return a.isTuple
}

func (a *Actual) ArgGOType() gotypes.Type {
	return a.argType
}

func (a *Actual) GetAsyncArg() *AsyncArg {
	return a.asyncArg
}

func (a *Actual) AppendWithArgsMethod() {
	if a.asyncArg.fun != nil {
		if len(a.asyncArg.fun.Args) > 0 {
			actualOrigFunc := a.Clone.Fun
			actualOrigArgs := a.Clone.Args

			actualOrigArgs[a.actualOffset] = a.asyncArg.fun.Fun
			call := &ast.SelectorExpr{
				Sel: ast.NewIdent("WithArguments"),
				X: &ast.CallExpr{
					Fun:  actualOrigFunc,
					Args: actualOrigArgs,
				},
			}

			a.Clone.Fun = call
			a.Clone.Args = a.asyncArg.fun.Args
			a.Clone = a.Clone.Fun.(*ast.SelectorExpr).X.(*ast.CallExpr)
		} else {
			a.Clone.Args[a.actualOffset] = a.asyncArg.fun.Fun
		}
	}
}

func (a *Actual) GetActualArg() ast.Expr {
	return a.Clone.Args[a.actualOffset]
}

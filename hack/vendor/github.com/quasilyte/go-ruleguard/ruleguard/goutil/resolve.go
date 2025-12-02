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

package goutil

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/ast/astutil"
)

func ResolveFunc(info *types.Info, callable ast.Expr) (ast.Expr, *types.Func) {
	switch callable := astutil.Unparen(callable).(type) {
	case *ast.Ident:
		sig, ok := info.ObjectOf(callable).(*types.Func)
		if !ok {
			return nil, nil
		}
		return nil, sig

	case *ast.SelectorExpr:
		sig, ok := info.ObjectOf(callable.Sel).(*types.Func)
		if !ok {
			return nil, nil
		}
		isMethod := sig.Type().(*types.Signature).Recv() != nil
		if _, ok := callable.X.(*ast.Ident); ok && !isMethod {
			return nil, sig
		}
		return callable.X, sig

	default:
		return nil, nil
	}
}

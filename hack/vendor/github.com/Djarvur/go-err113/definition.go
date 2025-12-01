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

package err113

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var methods2check = map[string]map[string]func(*ast.CallExpr, *types.Info) bool{ // nolint: gochecknoglobals
	"errors": {"New": justTrue},
	"fmt":    {"Errorf": checkWrap},
}

func justTrue(*ast.CallExpr, *types.Info) bool {
	return true
}

func checkWrap(ce *ast.CallExpr, info *types.Info) bool {
	return !(len(ce.Args) > 0 && strings.Contains(toString(ce.Args[0], info), `%w`))
}

func inspectDefinition(pass *analysis.Pass, tlds map[*ast.CallExpr]struct{}, n ast.Node) bool { //nolint: unparam
	// check whether the call expression matches time.Now().Sub()
	ce, ok := n.(*ast.CallExpr)
	if !ok {
		return true
	}

	if _, ok = tlds[ce]; ok {
		return true
	}

	fn, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok {
		return true
	}

	fxName, ok := asImportedName(fn.X, pass.TypesInfo)
	if !ok {
		return true
	}

	methods, ok := methods2check[fxName]
	if !ok {
		return true
	}

	checkFunc, ok := methods[fn.Sel.Name]
	if !ok {
		return true
	}

	if !checkFunc(ce, pass.TypesInfo) {
		return true
	}

	pass.Reportf(
		ce.Pos(),
		"do not define dynamic errors, use wrapped static errors instead: %q",
		render(pass.Fset, ce),
	)

	return true
}

func toString(ex ast.Expr, info *types.Info) string {
	if tv, ok := info.Types[ex]; ok && tv.Value != nil {
		return tv.Value.ExactString()
	}

	return ""
}

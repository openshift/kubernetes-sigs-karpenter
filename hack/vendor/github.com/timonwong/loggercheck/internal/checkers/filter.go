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

	"golang.org/x/tools/go/analysis"
)

func filterKeyAndValues(pass *analysis.Pass, keyAndValues []ast.Expr, objName string) []ast.Expr {
	// Check the argument count
	filtered := make([]ast.Expr, 0, len(keyAndValues))
	for _, arg := range keyAndValues {
		// Skip any object type field we found
		switch arg := arg.(type) {
		case *ast.CallExpr, *ast.Ident:
			typ := types.Unalias(pass.TypesInfo.TypeOf(arg))

			switch typ := typ.(type) {
			case *types.Named:
				obj := typ.Obj()
				if obj != nil && obj.Name() == objName {
					continue
				}

			default:
				// pass
			}
		}

		filtered = append(filtered, arg)
	}

	return filtered
}

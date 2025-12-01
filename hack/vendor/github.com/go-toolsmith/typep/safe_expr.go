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

package typep

import (
	"go/ast"
	"go/token"
	"go/types"
)

// SideEffectFree reports whether expr is softly safe expression and contains
// no significant side-effects. As opposed to strictly safe expressions,
// soft safe expressions permit some forms of side-effects, like
// panic possibility during indexing or nil pointer dereference.
//
// Uses types info to determine type conversion expressions that
// are the only permitted kinds of call expressions.
// Note that is does not check whether called function really
// has any side effects. The analysis is very conservative.
func SideEffectFree(info *types.Info, expr ast.Expr) bool {
	// This list switch is not comprehensive and uses
	// whitelist to be on the conservative side.
	// Can be extended as needed.

	if expr == nil {
		return true
	}

	switch expr := expr.(type) {
	case *ast.StarExpr:
		return SideEffectFree(info, expr.X)
	case *ast.BinaryExpr:
		return SideEffectFree(info, expr.X) &&
			SideEffectFree(info, expr.Y)
	case *ast.UnaryExpr:
		return expr.Op != token.ARROW &&
			SideEffectFree(info, expr.X)
	case *ast.BasicLit, *ast.Ident:
		return true
	case *ast.SliceExpr:
		return SideEffectFree(info, expr.X) &&
			SideEffectFree(info, expr.Low) &&
			SideEffectFree(info, expr.High) &&
			SideEffectFree(info, expr.Max)
	case *ast.IndexExpr:
		return SideEffectFree(info, expr.X) &&
			SideEffectFree(info, expr.Index)
	case *ast.SelectorExpr:
		return SideEffectFree(info, expr.X)
	case *ast.ParenExpr:
		return SideEffectFree(info, expr.X)
	case *ast.TypeAssertExpr:
		return SideEffectFree(info, expr.X)
	case *ast.CompositeLit:
		return SideEffectFreeList(info, expr.Elts)
	case *ast.CallExpr:
		return IsTypeExpr(info, expr.Fun) &&
			SideEffectFreeList(info, expr.Args)

	default:
		return false
	}
}

// SideEffectFreeList reports whether every expr in list is safe.
//
// See SideEffectFree.
func SideEffectFreeList(info *types.Info, list []ast.Expr) bool {
	for _, expr := range list {
		if !SideEffectFree(info, expr) {
			return false
		}
	}
	return true
}

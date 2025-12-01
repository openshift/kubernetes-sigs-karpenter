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

// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typesinternal

import (
	"go/ast"
	"go/token"
	"go/types"
)

// NoEffects reports whether the expression has no side effects, i.e., it
// does not modify the memory state. This function is conservative: it may
// return false even when the expression has no effect.
func NoEffects(info *types.Info, expr ast.Expr) bool {
	noEffects := true
	ast.Inspect(expr, func(n ast.Node) bool {
		switch v := n.(type) {
		case nil, *ast.Ident, *ast.BasicLit, *ast.BinaryExpr, *ast.ParenExpr,
			*ast.SelectorExpr, *ast.IndexExpr, *ast.SliceExpr, *ast.TypeAssertExpr,
			*ast.StarExpr, *ast.CompositeLit, *ast.ArrayType, *ast.StructType,
			*ast.MapType, *ast.InterfaceType, *ast.KeyValueExpr:
			// No effect
		case *ast.UnaryExpr:
			// Channel send <-ch has effects
			if v.Op == token.ARROW {
				noEffects = false
			}
		case *ast.CallExpr:
			// Type conversion has no effects
			if !info.Types[v.Fun].IsType() {
				// TODO(adonovan): Add a case for built-in functions without side
				// effects (by using callsPureBuiltin from tools/internal/refactor/inline)

				noEffects = false
			}
		case *ast.FuncLit:
			// A FuncLit has no effects, but do not descend into it.
			return false
		default:
			// All other expressions have effects
			noEffects = false
		}

		return noEffects
	})
	return noEffects
}

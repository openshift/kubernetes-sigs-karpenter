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

package analysisutil

import (
	"go/ast"
	"go/types"
)

// ObjectOf works in context of Golang package and returns types.Object for the given object's package and name.
// The search is based on the provided package and its dependencies (imports).
// Returns nil if the object is not found.
func ObjectOf(pkg *types.Package, objPkg, objName string) types.Object {
	if pkg.Path() == objPkg {
		return pkg.Scope().Lookup(objName)
	}

	for _, i := range pkg.Imports() {
		if trimVendor(i.Path()) == objPkg {
			return i.Scope().Lookup(objName)
		}
	}
	return nil
}

// IsObj returns true if expression is identifier which notes to given types.Object.
// Useful in combination with types.Universe objects.
func IsObj(typesInfo *types.Info, expr ast.Expr, expected types.Object) bool {
	id, ok := expr.(*ast.Ident)
	if !ok {
		return false
	}

	obj := typesInfo.ObjectOf(id)
	return obj == expected
}

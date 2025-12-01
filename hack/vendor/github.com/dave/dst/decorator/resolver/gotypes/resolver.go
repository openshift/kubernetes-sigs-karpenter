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

package gotypes

import (
	"errors"
	"go/ast"
	"go/types"
)

func New(uses map[*ast.Ident]types.Object) *DecoratorResolver {
	return &DecoratorResolver{Uses: uses}
}

type DecoratorResolver struct {
	Uses map[*ast.Ident]types.Object // Types info - must include Uses
}

func (r *DecoratorResolver) ResolveIdent(file *ast.File, parent ast.Node, parentField string, id *ast.Ident) (string, error) {

	if r.Uses == nil {
		return "", errors.New("gotypes.DecoratorResolver needs Uses in types info")
	}

	if se, ok := parent.(*ast.SelectorExpr); ok && parentField == "Sel" {

		// if the parent is a SelectorExpr and this Ident is in the Sel field, only resolve the path
		// if X is a package identifier

		xid, ok := se.X.(*ast.Ident)
		if !ok {
			// x is not an ident -> not a qualified identifier
			return "", nil
		}
		obj, ok := r.Uses[xid]
		if !ok {
			// not found in uses -> not a qualified identifier
			return "", nil
		}
		pn, ok := obj.(*types.PkgName)
		if !ok {
			// not a pkgname -> not a remote identifier
			return "", nil
		}
		return pn.Imported().Path(), nil
	}

	obj, ok := r.Uses[id]
	if !ok {
		// not found in uses -> not a remote identifier
		return "", nil
	}

	if v, ok := obj.(*types.Var); ok && v.IsField() {
		// field ident (e.g. name of a field in a composite literal) -> doesn't need qualified ident
		return "", nil
	}

	pkg := obj.Pkg()
	if pkg == nil {
		// pre-defined idents in the universe scope - e.g. "byte"
		return "", nil
	}

	return pkg.Path(), nil
}

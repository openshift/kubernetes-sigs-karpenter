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

// Package typeparams provides utilities for working with Go ASTs with support
// for type parameters when built with Go 1.18 and higher.
package typeparams

import (
	"go/ast"
)

// ReceiverType returns the named type of the method receiver, sans "*" and type
// parameters, or "invalid-type" if fn.Recv is ill formed.
func ReceiverType(fn *ast.FuncDecl) string {
	e := fn.Recv.List[0].Type
	if s, ok := e.(*ast.StarExpr); ok {
		e = s.X
	}
	e = unpackIndexExpr(e)
	if id, ok := e.(*ast.Ident); ok {
		return id.Name
	}
	return "invalid-type"
}

func unpackIndexExpr(e ast.Expr) ast.Expr {
	switch e := e.(type) {
	case *ast.IndexExpr:
		return e.X
	case *ast.IndexListExpr:
		return e.X
	}
	return e
}

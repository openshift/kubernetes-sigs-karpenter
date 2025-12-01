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

package astp

import "go/ast"

// IsDecl reports whether a node is a ast.Decl.
func IsDecl(node ast.Node) bool {
	_, ok := node.(ast.Decl)
	return ok
}

// IsFuncDecl reports whether a given ast.Node is a function declaration (*ast.FuncDecl).
func IsFuncDecl(node ast.Node) bool {
	_, ok := node.(*ast.FuncDecl)
	return ok
}

// IsGenDecl reports whether a given ast.Node is a generic declaration (*ast.GenDecl).
func IsGenDecl(node ast.Node) bool {
	_, ok := node.(*ast.GenDecl)
	return ok
}

// IsImportSpec reports whether a given ast.Node is an import declaration (*ast.ImportSpec).
func IsImportSpec(node ast.Node) bool {
	_, ok := node.(*ast.ImportSpec)
	return ok
}

// IsValueSpec reports whether a given ast.Node is a value declaration (*ast.ValueSpec).
func IsValueSpec(node ast.Node) bool {
	_, ok := node.(*ast.ValueSpec)
	return ok
}

// IsTypeSpec reports whether a given ast.Node is a type declaration (*ast.TypeSpec).
func IsTypeSpec(node ast.Node) bool {
	_, ok := node.(*ast.TypeSpec)
	return ok
}

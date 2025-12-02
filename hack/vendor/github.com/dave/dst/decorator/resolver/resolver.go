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

package resolver

import (
	"errors"
	"go/ast"
)

// RestorerResolver resolves a package path to a package name.
type RestorerResolver interface {
	ResolvePackage(path string) (string, error)
}

// DecoratorResolver resolves an identifier to a local or remote reference.
//
// Returns path == "" if the node is not a local or remote reference (e.g. a field in a composite
// literal, the selector in a selector expression etc.).
//
// Returns path == "" is the node is a local reference.
//
// Returns path != "" is the node is a remote reference.
type DecoratorResolver interface {
	ResolveIdent(file *ast.File, parent ast.Node, parentField string, id *ast.Ident) (path string, err error)
}

// ErrPackageNotFound means the package is not found
var ErrPackageNotFound = errors.New("package not found")

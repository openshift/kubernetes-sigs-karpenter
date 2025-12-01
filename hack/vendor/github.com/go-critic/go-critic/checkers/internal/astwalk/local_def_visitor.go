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

package astwalk

import (
	"go/ast"
)

// LocalDefVisitor visits every name definitions inside a function.
//
// Next elements are considered as name definitions:
//   - Function parameters (input, output, receiver)
//   - Every LHS of ":=" assignment that defines a new name
//   - Every local var/const declaration.
//
// NOTE: this visitor is experimental.
// This is also why it lives in a separate file.
type LocalDefVisitor interface {
	walkerEvents
	VisitLocalDef(Name, ast.Expr)
}

// NameKind describes what kind of name Name object holds.
type NameKind int

// Name holds ver/const/param definition symbol info.
type Name struct {
	ID   *ast.Ident
	Kind NameKind

	// Index is NameVar-specific field that is used to
	// specify nth tuple element being assigned to the name.
	Index int
}

// NOTE: set of name kinds is not stable and may change over time.
//
// TODO(quasilyte): is NameRecv/NameParam/NameResult granularity desired?
// TODO(quasilyte): is NameVar/NameBind (var vs :=) granularity desired?
const (
	// NameParam is function/method receiver/input/output name.
	// Initializing expression is always nil.
	NameParam NameKind = iota
	// NameVar is var or ":=" declared name.
	// Initializing expression may be nil for var-declared names
	// without explicit initializing expression.
	NameVar
	// NameConst is const-declared name.
	// Initializing expression is never nil.
	NameConst
)

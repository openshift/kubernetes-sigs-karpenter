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

package internal

import (
	"go/ast"
)

type StructConstructor struct {
	constructor  *ast.FuncDecl
	structReturn *ast.Ident
}

func NewStructConstructor(funcDec *ast.FuncDecl) (StructConstructor, bool) {
	if !FuncCanBeConstructor(funcDec) {
		return StructConstructor{}, false
	}

	expr := funcDec.Type.Results.List[0].Type

	returnType, ok := GetIdent(expr)
	if !ok {
		return StructConstructor{}, false
	}

	return StructConstructor{
		constructor:  funcDec,
		structReturn: returnType,
	}, true
}

// GetStructReturn Return the struct linked to this "constructor".
func (sc StructConstructor) GetStructReturn() *ast.Ident {
	return sc.structReturn
}

func (sc StructConstructor) GetConstructor() *ast.FuncDecl {
	return sc.constructor
}

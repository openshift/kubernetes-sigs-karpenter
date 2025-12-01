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

//go:build go1.18
// +build go1.18

package astcopy

import (
	"go/ast"
)

// FuncType returns x deep copy.
// Copy of nil argument is nil.
func FuncType(x *ast.FuncType) *ast.FuncType {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Params = FieldList(x.Params)
	cp.Results = FieldList(x.Results)
	cp.TypeParams = FieldList(x.TypeParams)
	return &cp
}

// TypeSpec returns x deep copy.
// Copy of nil argument is nil.
func TypeSpec(x *ast.TypeSpec) *ast.TypeSpec {
	if x == nil {
		return nil
	}
	cp := *x
	cp.Name = Ident(x.Name)
	cp.Type = copyExpr(x.Type)
	cp.Doc = CommentGroup(x.Doc)
	cp.Comment = CommentGroup(x.Comment)
	cp.TypeParams = FieldList(x.TypeParams)
	return &cp
}

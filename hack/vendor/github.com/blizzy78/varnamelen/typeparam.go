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

package varnamelen

import "go/ast"

// isTypeParam returns true if field is a type parameter of any of the given funcs.
func isTypeParam(field *ast.Field, funcs []*ast.FuncDecl, funcLits []*ast.FuncLit) bool { //nolint:gocognit // it's not that complicated
	for _, f := range funcs {
		if f.Type.TypeParams == nil {
			continue
		}

		for _, p := range f.Type.TypeParams.List {
			if p == field {
				return true
			}
		}
	}

	for _, f := range funcLits {
		if f.Type.TypeParams == nil {
			continue
		}

		for _, p := range f.Type.TypeParams.List {
			if p == field {
				return true
			}
		}
	}

	return false
}

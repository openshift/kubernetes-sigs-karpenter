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

// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astutil

import (
	"go/ast"
	"iter"
)

// FlatFields 'flattens' an ast.FieldList, returning an iterator over each
// (name, field) combination in the list. For unnamed fields, the identifier is
// nil.
func FlatFields(list *ast.FieldList) iter.Seq2[*ast.Ident, *ast.Field] {
	return func(yield func(*ast.Ident, *ast.Field) bool) {
		if list == nil {
			return
		}

		for _, field := range list.List {
			if len(field.Names) == 0 {
				if !yield(nil, field) {
					return
				}
			} else {
				for _, name := range field.Names {
					if !yield(name, field) {
						return
					}
				}
			}
		}
	}
}

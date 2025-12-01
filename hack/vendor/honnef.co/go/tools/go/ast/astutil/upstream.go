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

package astutil

import (
	"go/ast"
	"go/token"
	_ "unsafe"

	"golang.org/x/tools/go/ast/astutil"
)

type Cursor = astutil.Cursor
type ApplyFunc = astutil.ApplyFunc

func Apply(root ast.Node, pre, post ApplyFunc) (result ast.Node) {
	return astutil.Apply(root, pre, post)
}

func PathEnclosingInterval(root *ast.File, start, end token.Pos) (path []ast.Node, exact bool) {
	return astutil.PathEnclosingInterval(root, start, end)
}

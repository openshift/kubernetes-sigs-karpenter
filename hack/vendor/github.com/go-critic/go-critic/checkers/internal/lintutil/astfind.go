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

package lintutil

import (
	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

// FindNode applies pred for root and all it's childs until it returns true.
// If followFunc is defined, it's called before following any node to check whether it needs to be followed.
// followFunc has to return true in order to continuing traversing the node and return false otherwise.
// Matched node is returned.
// If none of the nodes matched predicate, nil is returned.
func FindNode(root ast.Node, followFunc, pred func(ast.Node) bool) ast.Node {
	var (
		found   ast.Node
		preFunc func(*astutil.Cursor) bool
	)

	if followFunc != nil {
		preFunc = func(cur *astutil.Cursor) bool {
			return followFunc(cur.Node())
		}
	}

	astutil.Apply(root,
		preFunc,
		func(cur *astutil.Cursor) bool {
			if pred(cur.Node()) {
				found = cur.Node()
				return false
			}
			return true
		})
	return found
}

// ContainsNode reports whether `FindNode(root, pred)!=nil`.
func ContainsNode(root ast.Node, pred func(ast.Node) bool) bool {
	return FindNode(root, nil, pred) != nil
}

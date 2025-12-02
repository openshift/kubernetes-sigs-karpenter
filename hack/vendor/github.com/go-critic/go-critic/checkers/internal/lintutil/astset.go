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

	"github.com/go-toolsmith/astequal"
)

// AstSet is a simple ast.Node set.
// Zero value is ready to use set.
// Can be reused after Clear call.
type AstSet struct {
	items []ast.Node
}

// Contains reports whether s contains x.
func (s *AstSet) Contains(x ast.Node) bool {
	for i := range s.items {
		if astequal.Node(s.items[i], x) {
			return true
		}
	}
	return false
}

// Insert pushes x in s if it's not already there.
// Returns true if element was inserted.
func (s *AstSet) Insert(x ast.Node) bool {
	if s.Contains(x) {
		return false
	}
	s.items = append(s.items, x)
	return true
}

// Clear removes all element from set.
func (s *AstSet) Clear() {
	s.items = s.items[:0]
}

// Len returns the number of elements contained inside s.
func (s *AstSet) Len() int {
	return len(s.items)
}

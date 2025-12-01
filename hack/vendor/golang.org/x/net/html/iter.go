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

//go:build go1.23

package html

import "iter"

// Ancestors returns an iterator over the ancestors of n, starting with n.Parent.
//
// Mutating a Node or its parents while iterating may have unexpected results.
func (n *Node) Ancestors() iter.Seq[*Node] {
	_ = n.Parent // eager nil check

	return func(yield func(*Node) bool) {
		for p := n.Parent; p != nil && yield(p); p = p.Parent {
		}
	}
}

// ChildNodes returns an iterator over the immediate children of n,
// starting with n.FirstChild.
//
// Mutating a Node or its children while iterating may have unexpected results.
func (n *Node) ChildNodes() iter.Seq[*Node] {
	_ = n.FirstChild // eager nil check

	return func(yield func(*Node) bool) {
		for c := n.FirstChild; c != nil && yield(c); c = c.NextSibling {
		}
	}

}

// Descendants returns an iterator over all nodes recursively beneath
// n, excluding n itself. Nodes are visited in depth-first preorder.
//
// Mutating a Node or its descendants while iterating may have unexpected results.
func (n *Node) Descendants() iter.Seq[*Node] {
	_ = n.FirstChild // eager nil check

	return func(yield func(*Node) bool) {
		n.descendants(yield)
	}
}

func (n *Node) descendants(yield func(*Node) bool) bool {
	for c := range n.ChildNodes() {
		if !yield(c) || !c.descendants(yield) {
			return false
		}
	}
	return true
}

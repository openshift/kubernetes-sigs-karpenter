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

package iradix

import "bytes"

// PathIterator is used to iterate over a set of nodes from the root
// down to a specified path. This will iterate over the same values that
// the Node.WalkPath method will.
type PathIterator[T any] struct {
	node *Node[T]
	path []byte
	done bool
}

// Next returns the next node in order
func (i *PathIterator[T]) Next() ([]byte, T, bool) {
	// This is mostly just an asynchronous implementation of the WalkPath
	// method on the node.
	var zero T
	var leaf *leafNode[T]

	for leaf == nil && i.node != nil {
		// visit the leaf values if any
		if i.node.leaf != nil {
			leaf = i.node.leaf
		}

		i.iterate()
	}

	if leaf != nil {
		return leaf.key, leaf.val, true
	}

	return nil, zero, false
}

func (i *PathIterator[T]) iterate() {
	// Check for key exhaustion
	if len(i.path) == 0 {
		i.node = nil
		return
	}

	// Look for an edge
	_, i.node = i.node.getEdge(i.path[0])
	if i.node == nil {
		return
	}

	// Consume the search prefix
	if bytes.HasPrefix(i.path, i.node.prefix) {
		i.path = i.path[len(i.node.prefix):]
	} else {
		// there are no more nodes to iterate through so
		// nil out the node to prevent returning results
		// for subsequent calls to Next()
		i.node = nil
	}
}

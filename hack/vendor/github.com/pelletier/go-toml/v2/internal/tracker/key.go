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

package tracker

import "github.com/pelletier/go-toml/v2/unstable"

// KeyTracker is a tracker that keeps track of the current Key as the AST is
// walked.
type KeyTracker struct {
	k []string
}

// UpdateTable sets the state of the tracker with the AST table node.
func (t *KeyTracker) UpdateTable(node *unstable.Node) {
	t.reset()
	t.Push(node)
}

// UpdateArrayTable sets the state of the tracker with the AST array table node.
func (t *KeyTracker) UpdateArrayTable(node *unstable.Node) {
	t.reset()
	t.Push(node)
}

// Push the given key on the stack.
func (t *KeyTracker) Push(node *unstable.Node) {
	it := node.Key()
	for it.Next() {
		t.k = append(t.k, string(it.Node().Data))
	}
}

// Pop key from stack.
func (t *KeyTracker) Pop(node *unstable.Node) {
	it := node.Key()
	for it.Next() {
		t.k = t.k[:len(t.k)-1]
	}
}

// Key returns the current key
func (t *KeyTracker) Key() []string {
	k := make([]string, len(t.k))
	copy(k, t.k)
	return k
}

func (t *KeyTracker) reset() {
	t.k = t.k[:0]
}

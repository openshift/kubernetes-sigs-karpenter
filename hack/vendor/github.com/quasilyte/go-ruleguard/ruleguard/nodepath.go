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

package ruleguard

import (
	"fmt"
	"go/ast"
	"strings"
)

type nodePath struct {
	stack []ast.Node
}

func newNodePath() *nodePath {
	return &nodePath{stack: make([]ast.Node, 0, 32)}
}

func (p nodePath) String() string {
	parts := make([]string, len(p.stack))
	for i, n := range p.stack {
		parts[i] = fmt.Sprintf("%T", n)
	}
	return strings.Join(parts, "/")
}

func (p *nodePath) Parent() ast.Node {
	return p.NthParent(1)
}

func (p *nodePath) Current() ast.Node {
	return p.NthParent(0)
}

func (p *nodePath) NthParent(n int) ast.Node {
	index := uint(len(p.stack) - n - 1)
	if index < uint(len(p.stack)) {
		return p.stack[index]
	}
	return nil
}

func (p *nodePath) Len() int { return len(p.stack) }

func (p *nodePath) Push(n ast.Node) {
	p.stack = append(p.stack, n)
}

func (p *nodePath) Pop() {
	p.stack = p.stack[:len(p.stack)-1]
}

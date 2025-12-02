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

package wsl

import (
	"go/ast"
)

// Cursor holds a list of statements and a pointer to where in the list we are.
// Each block gets a new cursor and can be used to check previous or coming
// statements.
type Cursor struct {
	currentIdx int
	statements []ast.Stmt
	checkType  CheckType
}

// NewCursor creates a new cursor with a given list of statements.
func NewCursor(statements []ast.Stmt) *Cursor {
	return &Cursor{
		currentIdx: -1,
		statements: statements,
	}
}

func (c *Cursor) SetChecker(ct CheckType) {
	c.checkType = ct
}

func (c *Cursor) NextNode() ast.Node {
	defer c.Save()()

	var nextNode ast.Node
	if c.Next() {
		nextNode = c.Stmt()
	}

	return nextNode
}

func (c *Cursor) Next() bool {
	if c.currentIdx >= len(c.statements)-1 {
		return false
	}

	c.currentIdx++

	return true
}

func (c *Cursor) Previous() bool {
	if c.currentIdx <= 0 {
		return false
	}

	c.currentIdx--

	return true
}

func (c *Cursor) PreviousNode() ast.Node {
	defer c.Save()()

	var previousNode ast.Node
	if c.Previous() {
		previousNode = c.Stmt()
	}

	return previousNode
}

func (c *Cursor) Stmt() ast.Stmt {
	return c.statements[c.currentIdx]
}

func (c *Cursor) Save() func() {
	idx := c.currentIdx

	return func() {
		c.currentIdx = idx
	}
}

func (c *Cursor) Len() int {
	return len(c.statements)
}

func (c *Cursor) Nth(n int) ast.Stmt {
	return c.statements[n]
}

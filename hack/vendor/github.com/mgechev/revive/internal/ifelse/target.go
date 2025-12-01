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

package ifelse

import "go/ast"

// Target decides what line/column should be indicated by the rule in question.
type Target int

const (
	// TargetIf means the text refers to the "if".
	TargetIf Target = iota

	// TargetElse means the text refers to the "else".
	TargetElse
)

func (t Target) node(ifStmt *ast.IfStmt) ast.Node {
	switch t {
	case TargetIf:
		return ifStmt
	case TargetElse:
		return ifStmt.Else
	}
	panic("bad target")
}

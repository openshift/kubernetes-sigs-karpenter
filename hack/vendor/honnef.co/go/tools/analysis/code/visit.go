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

package code

import (
	"bytes"
	"go/ast"
	"go/format"

	"honnef.co/go/tools/pattern"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func Preorder(pass *analysis.Pass, fn func(ast.Node), types ...ast.Node) {
	pass.ResultOf[inspect.Analyzer].(*inspector.Inspector).Preorder(types, fn)
}

func PreorderStack(pass *analysis.Pass, fn func(ast.Node, []ast.Node), types ...ast.Node) {
	pass.ResultOf[inspect.Analyzer].(*inspector.Inspector).WithStack(types, func(n ast.Node, push bool, stack []ast.Node) (proceed bool) {
		if push {
			fn(n, stack)
		}
		return true
	})
}

func Match(pass *analysis.Pass, q pattern.Pattern, node ast.Node) (*pattern.Matcher, bool) {
	// Note that we ignore q.Relevant – callers of Match usually use
	// AST inspectors that already filter on nodes we're interested
	// in.
	m := &pattern.Matcher{TypesInfo: pass.TypesInfo}
	ok := m.Match(q, node)
	return m, ok
}

func MatchAndEdit(pass *analysis.Pass, before, after pattern.Pattern, node ast.Node) (*pattern.Matcher, []analysis.TextEdit, bool) {
	m, ok := Match(pass, before, node)
	if !ok {
		return m, nil, false
	}
	r := pattern.NodeToAST(after.Root, m.State)
	buf := &bytes.Buffer{}
	format.Node(buf, pass.Fset, r)
	edit := []analysis.TextEdit{{
		Pos:     node.Pos(),
		End:     node.End(),
		NewText: buf.Bytes(),
	}}
	return m, edit, true
}

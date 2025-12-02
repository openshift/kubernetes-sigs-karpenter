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

package rule

import (
	"go/ast"
	"go/types"

	"github.com/mgechev/revive/lint"
)

// StringOfIntRule warns when logic expressions contains Boolean literals.
type StringOfIntRule struct{}

// Apply applies the rule to given file.
func (*StringOfIntRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}

	astFile := file.AST
	file.Pkg.TypeCheck()

	w := &lintStringInt{file, onFailure}
	ast.Walk(w, astFile)

	return failures
}

// Name returns the rule name.
func (*StringOfIntRule) Name() string {
	return "string-of-int"
}

type lintStringInt struct {
	file      *lint.File
	onFailure func(lint.Failure)
}

func (w *lintStringInt) Visit(node ast.Node) ast.Visitor {
	ce, ok := node.(*ast.CallExpr)
	if !ok {
		return w
	}

	if !w.isCallStringCast(ce.Fun) {
		return w
	}

	if !w.isIntExpression(ce.Args) {
		return w
	}

	w.onFailure(lint.Failure{
		Confidence: 1,
		Node:       ce,
		Failure:    "dubious conversion of an integer into a string, use strconv.Itoa",
	})

	return w
}

func (w *lintStringInt) isCallStringCast(e ast.Expr) bool {
	t := w.file.Pkg.TypeOf(e)
	if t == nil {
		return false
	}

	tb, _ := t.Underlying().(*types.Basic)

	return tb != nil && tb.Kind() == types.String
}

func (w *lintStringInt) isIntExpression(es []ast.Expr) bool {
	if len(es) != 1 {
		return false
	}

	t := w.file.Pkg.TypeOf(es[0])
	if t == nil {
		return false
	}

	ut, _ := t.Underlying().(*types.Basic)
	if ut == nil || ut.Info()&types.IsInteger == 0 {
		return false
	}

	switch ut.Kind() {
	case types.Byte, types.Rune, types.UntypedRune:
		return false
	}

	return true
}

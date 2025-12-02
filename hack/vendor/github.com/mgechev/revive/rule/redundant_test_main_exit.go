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
	"fmt"
	"go/ast"

	"github.com/mgechev/revive/lint"
)

// RedundantTestMainExitRule suggests removing Exit call in TestMain function for test files.
type RedundantTestMainExitRule struct{}

// Apply applies the rule to given file.
func (*RedundantTestMainExitRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	if !file.IsTest() || !file.Pkg.IsAtLeastGoVersion(lint.Go115) {
		// skip analysis for non-test files or for Go versions before 1.15
		return failures
	}

	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}

	w := &lintRedundantTestMainExit{onFailure: onFailure}
	ast.Walk(w, file.AST)
	return failures
}

// Name returns the rule name.
func (*RedundantTestMainExitRule) Name() string {
	return "redundant-test-main-exit"
}

type lintRedundantTestMainExit struct {
	onFailure func(lint.Failure)
}

func (w *lintRedundantTestMainExit) Visit(node ast.Node) ast.Visitor {
	if fd, ok := node.(*ast.FuncDecl); ok {
		if fd.Name.Name != "TestMain" {
			return nil // skip analysis for other functions than TestMain
		}

		return w
	}

	se, ok := node.(*ast.ExprStmt)
	if !ok {
		return w
	}
	ce, ok := se.X.(*ast.CallExpr)
	if !ok {
		return w
	}

	fc, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok {
		return w
	}
	id, ok := fc.X.(*ast.Ident)
	if !ok {
		return w
	}

	pkg := id.Name
	fn := fc.Sel.Name
	if isCallToExitFunction(pkg, fn) {
		w.onFailure(lint.Failure{
			Confidence: 1,
			Node:       ce,
			Category:   lint.FailureCategoryStyle,
			Failure:    fmt.Sprintf("redundant call to %s.%s in TestMain function, the test runner will handle it automatically as of Go 1.15", pkg, fn),
		})
	}

	return w
}

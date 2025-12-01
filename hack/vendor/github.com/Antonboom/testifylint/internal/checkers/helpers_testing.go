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

package checkers

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func isSubTestRun(pass *analysis.Pass, ce *ast.CallExpr) bool {
	se, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok || se.Sel == nil {
		return false
	}
	return (implementsTestingT(pass, se.X) || implementsTestifySuite(pass, se.X)) && se.Sel.Name == "Run"
}

func isTestingFuncOrMethod(pass *analysis.Pass, fd *ast.FuncDecl) bool {
	return hasTestingTParam(pass, fd.Type) || isSuiteMethod(pass, fd)
}

func isTestingAnonymousFunc(pass *analysis.Pass, ft *ast.FuncType) bool {
	return hasTestingTParam(pass, ft)
}

func hasTestingTParam(pass *analysis.Pass, ft *ast.FuncType) bool {
	if ft == nil || ft.Params == nil {
		return false
	}

	for _, param := range ft.Params.List {
		if implementsTestingT(pass, param.Type) {
			return true
		}
	}
	return false
}

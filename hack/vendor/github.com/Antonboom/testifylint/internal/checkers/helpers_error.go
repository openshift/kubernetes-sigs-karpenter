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
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var (
	errorObj   = types.Universe.Lookup("error")
	errorType  = errorObj.Type()
	errorIface = errorType.Underlying().(*types.Interface)
)

func isError(pass *analysis.Pass, expr ast.Expr) bool {
	return pass.TypesInfo.TypeOf(expr) == errorType
}

func isErrorsIsCall(pass *analysis.Pass, ce *ast.CallExpr) bool {
	return isPkgFnCall(pass, ce, "errors", "Is")
}

func isErrorsAsCall(pass *analysis.Pass, ce *ast.CallExpr) bool {
	return isPkgFnCall(pass, ce, "errors", "As")
}

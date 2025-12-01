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

	"github.com/Antonboom/testifylint/internal/analysisutil"
)

var lenObj = types.Universe.Lookup("len")

func isLenCallAndZero(pass *analysis.Pass, a, b ast.Expr) (ast.Expr, bool) {
	lenArg, ok := isBuiltinLenCall(pass, a)
	return lenArg, ok && isZero(b)
}

func isBuiltinLenCall(pass *analysis.Pass, e ast.Expr) (ast.Expr, bool) {
	ce, ok := e.(*ast.CallExpr)
	if !ok {
		return nil, false
	}

	if analysisutil.IsObj(pass.TypesInfo, ce.Fun, lenObj) && len(ce.Args) == 1 {
		return ce.Args[0], true
	}
	return nil, false
}

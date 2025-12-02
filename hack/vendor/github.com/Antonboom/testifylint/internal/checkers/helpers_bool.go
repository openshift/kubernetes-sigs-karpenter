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

var (
	falseObj = types.Universe.Lookup("false")
	trueObj  = types.Universe.Lookup("true")
)

func isUntypedBool(pass *analysis.Pass, e ast.Expr) bool {
	return isUntypedTrue(pass, e) || isUntypedFalse(pass, e)
}

func isUntypedTrue(pass *analysis.Pass, e ast.Expr) bool {
	return analysisutil.IsObj(pass.TypesInfo, e, trueObj)
}

func isUntypedFalse(pass *analysis.Pass, e ast.Expr) bool {
	return analysisutil.IsObj(pass.TypesInfo, e, falseObj)
}

func hasBoolType(pass *analysis.Pass, e ast.Expr) bool {
	basicType, ok := pass.TypesInfo.TypeOf(e).(*types.Basic)
	return ok && basicType.Kind() == types.Bool
}

func isBoolOverride(pass *analysis.Pass, e ast.Expr) bool {
	namedType, ok := pass.TypesInfo.TypeOf(e).(*types.Named)
	return ok && namedType.Obj().Name() == "bool"
}

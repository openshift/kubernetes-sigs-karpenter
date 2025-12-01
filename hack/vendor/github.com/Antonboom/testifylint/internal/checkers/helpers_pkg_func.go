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

	"github.com/Antonboom/testifylint/internal/analysisutil"
)

func isFmtSprintfCall(pass *analysis.Pass, e ast.Expr) ([]ast.Expr, bool) {
	ce, ok := e.(*ast.CallExpr)
	if !ok {
		return nil, false
	}
	return ce.Args, isPkgFnCall(pass, ce, "fmt", "Sprintf")
}

func isJSONRawMessageCast(pass *analysis.Pass, ce *ast.CallExpr) bool {
	return isPkgFnCall(pass, ce, "encoding/json", "RawMessage")
}

func isRegexpMustCompileCall(pass *analysis.Pass, ce *ast.CallExpr) bool {
	return isPkgFnCall(pass, ce, "regexp", "MustCompile")
}

func isStringsContainsCall(pass *analysis.Pass, ce *ast.CallExpr) bool {
	return isPkgFnCall(pass, ce, "strings", "Contains")
}

func isStringsReplaceCall(pass *analysis.Pass, ce *ast.CallExpr) bool {
	return isPkgFnCall(pass, ce, "strings", "Replace")
}

func isStringsReplaceAllCall(pass *analysis.Pass, ce *ast.CallExpr) bool {
	return isPkgFnCall(pass, ce, "strings", "ReplaceAll")
}

func isStringsTrimCall(pass *analysis.Pass, ce *ast.CallExpr) bool {
	return isPkgFnCall(pass, ce, "strings", "Trim")
}

func isStringsTrimSpaceCall(pass *analysis.Pass, ce *ast.CallExpr) bool {
	return isPkgFnCall(pass, ce, "strings", "TrimSpace")
}

func isPkgFnCall(pass *analysis.Pass, ce *ast.CallExpr, pkg, fn string) bool {
	se, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	fnObj := analysisutil.ObjectOf(pass.Pkg, pkg, fn)
	if fnObj == nil {
		return false
	}

	return analysisutil.IsObj(pass.TypesInfo, se.Sel, fnObj)
}

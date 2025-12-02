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

type Config struct {
	RequireStringKey bool
	NoPrintfLike     bool
}

type CallContext struct {
	Expr      *ast.CallExpr
	Func      *types.Func
	Signature *types.Signature
}

type Checker interface {
	FilterKeyAndValues(pass *analysis.Pass, keyAndValues []ast.Expr) []ast.Expr
	CheckLoggingKey(pass *analysis.Pass, keyAndValues []ast.Expr)
	CheckPrintfLikeSpecifier(pass *analysis.Pass, args []ast.Expr)
}

func ExecuteChecker(c Checker, pass *analysis.Pass, call CallContext, cfg Config) {
	params := call.Signature.Params()
	nparams := params.Len() // variadic => nonzero
	startIndex := nparams - 1

	iface, ok := types.Unalias(params.At(startIndex).Type().(*types.Slice).Elem()).(*types.Interface)
	if !ok || !iface.Empty() {
		return // final (args) param is not ...interface{}
	}

	keyValuesArgs := c.FilterKeyAndValues(pass, call.Expr.Args[startIndex:])

	if len(keyValuesArgs)%2 != 0 {
		firstArg := keyValuesArgs[0]
		lastArg := keyValuesArgs[len(keyValuesArgs)-1]
		pass.Report(analysis.Diagnostic{
			Pos:      firstArg.Pos(),
			End:      lastArg.End(),
			Category: DiagnosticCategory,
			Message:  "odd number of arguments passed as key-value pairs for logging",
		})
	}

	if cfg.RequireStringKey {
		c.CheckLoggingKey(pass, keyValuesArgs)
	}

	if cfg.NoPrintfLike {
		// Check all args
		c.CheckPrintfLikeSpecifier(pass, call.Expr.Args)
	}
}

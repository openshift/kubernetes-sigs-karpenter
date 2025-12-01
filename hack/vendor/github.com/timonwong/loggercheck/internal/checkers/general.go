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
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/timonwong/loggercheck/internal/checkers/printf"
	"github.com/timonwong/loggercheck/internal/stringutil"
)

type General struct{}

func (g General) FilterKeyAndValues(_ *analysis.Pass, keyAndValues []ast.Expr) []ast.Expr {
	return keyAndValues
}

func (g General) CheckLoggingKey(pass *analysis.Pass, keyAndValues []ast.Expr) {
	for i := 0; i < len(keyAndValues); i += 2 {
		arg := keyAndValues[i]
		if value, ok := extractValueFromStringArg(pass, arg); ok {
			if stringutil.IsASCII(value) {
				continue
			}

			pass.Report(analysis.Diagnostic{
				Pos:      arg.Pos(),
				End:      arg.End(),
				Category: DiagnosticCategory,
				Message: fmt.Sprintf(
					"logging keys are expected to be alphanumeric strings, please remove any non-latin characters from %q",
					value),
			})
		} else {
			pass.Report(analysis.Diagnostic{
				Pos:      arg.Pos(),
				End:      arg.End(),
				Category: DiagnosticCategory,
				Message: fmt.Sprintf(
					"logging keys are expected to be inlined constant strings, please replace %q provided with string",
					renderNodeEllipsis(pass.Fset, arg)),
			})
		}
	}
}

func (g General) CheckPrintfLikeSpecifier(pass *analysis.Pass, args []ast.Expr) {
	for _, arg := range args {
		format, ok := extractValueFromStringArg(pass, arg)
		if !ok {
			continue
		}

		if specifier, ok := printf.IsPrintfLike(format); ok {
			pass.Report(analysis.Diagnostic{
				Pos:      arg.Pos(),
				End:      arg.End(),
				Category: DiagnosticCategory,
				Message:  fmt.Sprintf("logging message should not use format specifier %q", specifier),
			})

			return // One error diagnostic is enough
		}
	}
}

var _ Checker = (*General)(nil)

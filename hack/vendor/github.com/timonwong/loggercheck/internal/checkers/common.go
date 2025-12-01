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
	"go/constant"
	"go/printer"
	"go/token"
	"go/types"
	"strings"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

const (
	DiagnosticCategory = "logging"
)

// extractValueFromStringArg returns true if the argument is a string type (literal or constant).
func extractValueFromStringArg(pass *analysis.Pass, arg ast.Expr) (value string, ok bool) {
	if typeAndValue, ok := pass.TypesInfo.Types[arg]; ok {
		if typ, ok := typeAndValue.Type.(*types.Basic); ok && typ.Kind() == types.String && typeAndValue.Value != nil {
			return constant.StringVal(typeAndValue.Value), true
		}
	}

	return "", false
}

func renderNodeEllipsis(fset *token.FileSet, v interface{}) string {
	const maxLen = 20

	buf := &strings.Builder{}
	_ = printer.Fprint(buf, fset, v)
	s := buf.String()
	if utf8.RuneCountInString(s) > maxLen {
		// Copied from go/constant/value.go
		i := 0
		for n := 0; n < maxLen-3; n++ {
			_, size := utf8.DecodeRuneInString(s[i:])
			i += size
		}
		s = s[:i] + "..."
	}
	return s
}

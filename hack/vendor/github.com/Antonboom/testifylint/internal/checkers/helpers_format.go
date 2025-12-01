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
	"bytes"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/Antonboom/testifylint/internal/analysisutil"
)

// formatAsCallArgs joins a, b and c and returns bytes like `a, b, c`.
func formatAsCallArgs(pass *analysis.Pass, args ...ast.Expr) []byte {
	if len(args) == 0 {
		return []byte("")
	}

	var buf bytes.Buffer
	for i, arg := range args {
		buf.Write(analysisutil.NodeBytes(pass.Fset, arg))
		if i != len(args)-1 {
			buf.WriteString(", ")
		}
	}
	return buf.Bytes()
}

func formatWithStringCastForBytes(pass *analysis.Pass, e ast.Expr) []byte {
	if !hasBytesType(pass, e) {
		return analysisutil.NodeBytes(pass.Fset, e)
	}

	if se, ok := isBufferBytesCall(pass, e); ok {
		return []byte(analysisutil.NodeString(pass.Fset, se) + ".String()")
	}
	return []byte("string(" + analysisutil.NodeString(pass.Fset, e) + ")")
}

func isBufferBytesCall(pass *analysis.Pass, e ast.Expr) (ast.Node, bool) {
	ce, ok := e.(*ast.CallExpr)
	if !ok {
		return nil, false
	}

	se, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, false
	}

	if !isIdentWithName("Bytes", se.Sel) {
		return nil, false
	}
	if t := pass.TypesInfo.TypeOf(se.X); t != nil {
		// NOTE(a.telyshev): This is hack, because `bytes` package can be not imported,
		// and we cannot do "true" comparison with `Buffer` object.
		return se.X, strings.TrimPrefix(t.String(), "*") == "bytes.Buffer"
	}

	return nil, false
}

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

// Contains detects situations like
//
//	assert.True(t, strings.Contains(a, "abc123"))
//	assert.False(t, !strings.Contains(a, "abc123"))
//
//	assert.False(t, strings.Contains(a, "abc123"))
//	assert.True(t, !strings.Contains(a, "abc123"))
//
// and requires
//
//	assert.Contains(t, a, "abc123")
//	assert.NotContains(t, a, "abc123")
type Contains struct{}

// NewContains constructs Contains checker.
func NewContains() Contains   { return Contains{} }
func (Contains) Name() string { return "contains" }

func (checker Contains) Check(pass *analysis.Pass, call *CallMeta) *analysis.Diagnostic {
	if len(call.Args) < 1 {
		return nil
	}

	expr := call.Args[0]
	unpacked, isNeg := isNegation(expr)
	if isNeg {
		expr = unpacked
	}

	ce, ok := expr.(*ast.CallExpr)
	if !ok || len(ce.Args) != 2 {
		return nil
	}

	if !isStringsContainsCall(pass, ce) {
		return nil
	}

	var proposed string
	switch call.Fn.NameFTrimmed {
	default:
		return nil

	case "True":
		proposed = "Contains"
		if isNeg {
			proposed = "NotContains"
		}

	case "False":
		proposed = "NotContains"
		if isNeg {
			proposed = "Contains"
		}
	}

	return newUseFunctionDiagnostic(checker.Name(), call, proposed,
		analysis.TextEdit{
			Pos:     call.Args[0].Pos(),
			End:     call.Args[0].End(),
			NewText: formatAsCallArgs(pass, ce.Args[0], ce.Args[1]),
		})
}

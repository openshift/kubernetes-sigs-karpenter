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

// Regexp detects situations like
//
//	assert.Regexp(t, regexp.MustCompile(`\[.*\] DEBUG \(.*TestNew.*\): message`), out)
//	assert.NotRegexp(t, regexp.MustCompile(`\[.*\] TRACE message`), out)
//
// and requires
//
//	assert.Regexp(t, `\[.*\] DEBUG \(.*TestNew.*\): message`, out)
//	assert.NotRegexp(t, `\[.*\] TRACE message`, out)
type Regexp struct{}

// NewRegexp constructs Regexp checker.
func NewRegexp() Regexp     { return Regexp{} }
func (Regexp) Name() string { return "regexp" }

func (checker Regexp) Check(pass *analysis.Pass, call *CallMeta) *analysis.Diagnostic {
	switch call.Fn.NameFTrimmed {
	default:
		return nil
	case "Regexp", "NotRegexp":
	}

	if len(call.Args) < 1 {
		return nil
	}

	ce, ok := call.Args[0].(*ast.CallExpr)
	if !ok || len(ce.Args) != 1 {
		return nil
	}

	if isRegexpMustCompileCall(pass, ce) {
		return newRemoveMustCompileDiagnostic(pass, checker.Name(), call, ce, ce.Args[0])
	}
	return nil
}

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
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// EqualValues detects situations like
//
//	assert.EqualValues(t, 42, result.IntField)
//	assert.NotEqualValues(t, 42, result.IntField)
//	...
//
// and requires
//
//	assert.Equal(t, 42, result.IntField)
//	assert.NotEqual(t, 42, result.IntField)
type EqualValues struct{}

// NewEqualValues constructs EqualValues checker.
func NewEqualValues() EqualValues { return EqualValues{} }
func (EqualValues) Name() string  { return "equal-values" }

func (checker EqualValues) Check(pass *analysis.Pass, call *CallMeta) *analysis.Diagnostic {
	assrn := call.Fn.NameFTrimmed
	switch assrn {
	default:
		return nil
	case "EqualValues", "NotEqualValues":
	}

	if len(call.Args) < 2 {
		return nil
	}
	first, second := call.Args[0], call.Args[1]

	if isFunc(pass, first) || isFunc(pass, second) {
		// NOTE(a.telyshev): EqualValues for funcs is ok, but not Equal:
		// https://github.com/stretchr/testify/issues/1524
		return nil
	}

	ft, st := pass.TypesInfo.TypeOf(first), pass.TypesInfo.TypeOf(second)
	if !types.Identical(ft, st) {
		return nil
	}

	// Type of one of arguments is equivalent to any.
	if isEmptyInterfaceType(ft) || isEmptyInterfaceType(st) {
		// EqualValues is ok here.
		// Equal would check their types and would fail.
		return nil
	}

	proposed := strings.TrimSuffix(assrn, "Values")
	return newUseFunctionDiagnostic(checker.Name(), call, proposed)
}

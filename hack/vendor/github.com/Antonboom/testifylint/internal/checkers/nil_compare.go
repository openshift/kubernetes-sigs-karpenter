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
	"golang.org/x/tools/go/analysis"

	"github.com/Antonboom/testifylint/internal/analysisutil"
)

// NilCompare detects situations like
//
//	assert.Equal(t, nil, value)
//	assert.EqualValues(t, nil, value)
//	assert.Exactly(t, nil, value)
//
//	assert.NotEqual(t, nil, value)
//	assert.NotEqualValues(t, nil, value)
//
// and requires
//
//	assert.Nil(t, value)
//	assert.NotNil(t, value)
type NilCompare struct{}

// NewNilCompare constructs NilCompare checker.
func NewNilCompare() NilCompare { return NilCompare{} }
func (NilCompare) Name() string { return "nil-compare" }

func (checker NilCompare) Check(pass *analysis.Pass, call *CallMeta) *analysis.Diagnostic {
	if len(call.Args) < 2 {
		return nil
	}

	survivingArg, ok := xorNil(call.Args[0], call.Args[1])
	if !ok {
		return nil
	}

	var proposedFn string

	switch call.Fn.NameFTrimmed {
	case "Equal", "EqualValues", "Exactly":
		proposedFn = "Nil"
	case "NotEqual", "NotEqualValues":
		proposedFn = "NotNil"
	default:
		return nil
	}

	return newUseFunctionDiagnostic(checker.Name(), call, proposedFn,
		analysis.TextEdit{
			Pos:     call.Args[0].Pos(),
			End:     call.Args[1].End(),
			NewText: analysisutil.NodeBytes(pass.Fset, survivingArg),
		})
}

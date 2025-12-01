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

package sa4015

import (
	"fmt"
	"go/types"

	"honnef.co/go/tools/analysis/callcheck"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/go/ir"
	"honnef.co/go/tools/go/ir/irutil"
	"honnef.co/go/tools/go/types/typeutil"
	"honnef.co/go/tools/internal/passes/buildir"

	"golang.org/x/tools/go/analysis"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA4015",
		Requires: []*analysis.Analyzer{buildir.Analyzer},
		Run:      callcheck.Analyzer(checkMathIntRules),
	},
	Doc: &lint.RawDocumentation{
		Title:    `Calling functions like \'math.Ceil\' on floats converted from integers doesn't do anything useful`,
		Since:    "2017.1",
		Severity: lint.SeverityWarning,
		MergeIf:  lint.MergeIfAll,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var checkMathIntRules = map[string]callcheck.Check{
	"math.Ceil":  pointlessIntMath,
	"math.Floor": pointlessIntMath,
	"math.IsNaN": pointlessIntMath,
	"math.Trunc": pointlessIntMath,
	"math.IsInf": pointlessIntMath,
}

func pointlessIntMath(call *callcheck.Call) {
	if ConvertedFromInt(call.Args[0].Value) {
		call.Invalid(fmt.Sprintf("calling %s on a converted integer is pointless", irutil.CallName(call.Instr.Common())))
	}
}

func ConvertedFromInt(v callcheck.Value) bool {
	conv, ok := v.Value.(*ir.Convert)
	if !ok {
		return false
	}
	return typeutil.NewTypeSet(conv.X.Type()).All(func(t *types.Term) bool {
		b, ok := t.Type().Underlying().(*types.Basic)
		return ok && b.Info()&types.IsInteger != 0
	})
}

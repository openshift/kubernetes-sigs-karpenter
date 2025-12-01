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

package sa1018

import (
	"fmt"
	"go/constant"

	"honnef.co/go/tools/analysis/callcheck"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/go/ir"
	"honnef.co/go/tools/internal/passes/buildir"

	"golang.org/x/tools/go/analysis"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA1018",
		Requires: []*analysis.Analyzer{buildir.Analyzer},
		Run:      callcheck.Analyzer(rules),
	},
	Doc: &lint.RawDocumentation{
		Title: `\'strings.Replace\' called with \'n == 0\', which does nothing`,
		Text: `With \'n == 0\', zero instances will be replaced. To replace all
instances, use a negative number, or use \'strings.ReplaceAll\'.`,
		Since:    "2017.1",
		Severity: lint.SeverityWarning,
		MergeIf:  lint.MergeIfAny, // MergeIfAny if we only flag literals, not named constants
	},
})

var Analyzer = SCAnalyzer.Analyzer

var rules = map[string]callcheck.Check{
	"strings.Replace": check("strings.Replace", 3),
	"bytes.Replace":   check("bytes.Replace", 3),
}

func check(name string, arg int) callcheck.Check {
	return func(call *callcheck.Call) {
		arg := call.Args[arg]
		if k, ok := arg.Value.Value.(*ir.Const); ok && k.Value.Kind() == constant.Int {
			if v, ok := constant.Int64Val(k.Value); ok && v == 0 {
				arg.Invalid(fmt.Sprintf("calling %s with n == 0 will return no results, did you mean -1?", name))
			}
		}
	}
}

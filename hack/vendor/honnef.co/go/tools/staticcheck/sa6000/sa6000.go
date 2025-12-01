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

package sa6000

import (
	"fmt"

	"honnef.co/go/tools/analysis/callcheck"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/go/ir"
	"honnef.co/go/tools/go/ir/irutil"
	"honnef.co/go/tools/internal/passes/buildir"

	"golang.org/x/tools/go/analysis"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA6000",
		Requires: []*analysis.Analyzer{buildir.Analyzer},
		Run:      callcheck.Analyzer(rules),
	},
	Doc: &lint.RawDocumentation{
		Title:    `Using \'regexp.Match\' or related in a loop, should use \'regexp.Compile\'`,
		Since:    "2017.1",
		Severity: lint.SeverityWarning,
		MergeIf:  lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var rules = map[string]callcheck.Check{
	"regexp.Match":       check("regexp.Match"),
	"regexp.MatchReader": check("regexp.MatchReader"),
	"regexp.MatchString": check("regexp.MatchString"),
}

func check(name string) callcheck.Check {
	return func(call *callcheck.Call) {
		if callcheck.ExtractConst(call.Args[0].Value) == nil {
			return
		}
		if !isInLoop(call.Instr.Block()) {
			return
		}
		call.Invalid(fmt.Sprintf("calling %s in a loop has poor performance, consider using regexp.Compile", name))
	}
}

func isInLoop(b *ir.BasicBlock) bool {
	sets := irutil.FindLoops(b.Parent())
	for _, set := range sets {
		if set.Has(b) {
			return true
		}
	}
	return false
}

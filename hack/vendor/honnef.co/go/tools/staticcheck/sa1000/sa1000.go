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

package sa1000

import (
	"go/constant"
	"regexp"

	"honnef.co/go/tools/analysis/callcheck"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/internal/passes/buildir"

	"golang.org/x/tools/go/analysis"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA1000",
		Requires: []*analysis.Analyzer{buildir.Analyzer},
		Run:      callcheck.Analyzer(rules),
	},
	Doc: &lint.RawDocumentation{
		Title:    `Invalid regular expression`,
		Since:    "2017.1",
		Severity: lint.SeverityError,
		MergeIf:  lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var rules = map[string]callcheck.Check{
	"regexp.MustCompile": check,
	"regexp.Compile":     check,
	"regexp.Match":       check,
	"regexp.MatchReader": check,
	"regexp.MatchString": check,
}

func check(call *callcheck.Call) {
	arg := call.Args[0]
	if c := callcheck.ExtractConstExpectKind(arg.Value, constant.String); c != nil {
		s := constant.StringVal(c.Value)
		if _, err := regexp.Compile(s); err != nil {
			arg.Invalid(err.Error())
		}
	}
}

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

package sa1007

import (
	"fmt"
	"go/constant"
	"net/url"

	"honnef.co/go/tools/analysis/callcheck"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/internal/passes/buildir"
	"honnef.co/go/tools/knowledge"

	"golang.org/x/tools/go/analysis"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA1007",
		Requires: []*analysis.Analyzer{buildir.Analyzer},
		Run:      callcheck.Analyzer(rules),
	},
	Doc: &lint.RawDocumentation{
		Title:    `Invalid URL in \'net/url.Parse\'`,
		Since:    "2017.1",
		Severity: lint.SeverityError,
		MergeIf:  lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var rules = map[string]callcheck.Check{
	"net/url.Parse": func(call *callcheck.Call) {
		arg := call.Args[knowledge.Arg("net/url.Parse.rawurl")]
		if c := callcheck.ExtractConstExpectKind(arg.Value, constant.String); c != nil {
			s := constant.StringVal(c.Value)
			_, err := url.Parse(s)
			if err != nil {
				arg.Invalid(fmt.Sprintf("%q is not a valid URL: %s", s, err))
			}
		}
	},
}

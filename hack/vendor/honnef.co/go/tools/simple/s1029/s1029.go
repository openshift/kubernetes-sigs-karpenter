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

package s1029

import (
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/internal/passes/buildir"
	"honnef.co/go/tools/internal/sharedcheck"

	"golang.org/x/tools/go/analysis"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "S1029",
		Run:      sharedcheck.CheckRangeStringRunes,
		Requires: []*analysis.Analyzer{buildir.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title: `Range over the string directly`,
		Text: `Ranging over a string will yield byte offsets and runes. If the offset
isn't used, this is functionally equivalent to converting the string
to a slice of runes and ranging over that. Ranging directly over the
string will be more performant, however, as it avoids allocating a new
slice, the size of which depends on the length of the string.`,
		Before:  `for _, r := range []rune(s) {}`,
		After:   `for _, r := range s {}`,
		Since:   "2017.1",
		MergeIf: lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

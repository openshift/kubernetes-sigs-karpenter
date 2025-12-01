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

package sa6003

import (
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/internal/passes/buildir"
	"honnef.co/go/tools/internal/sharedcheck"

	"golang.org/x/tools/go/analysis"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA6003",
		Run:      sharedcheck.CheckRangeStringRunes,
		Requires: []*analysis.Analyzer{buildir.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title: `Converting a string to a slice of runes before ranging over it`,
		Text: `You may want to loop over the runes in a string. Instead of converting
the string to a slice of runes and looping over that, you can loop
over the string itself. That is,

    for _, r := range s {}

and

    for _, r := range []rune(s) {}

will yield the same values. The first version, however, will be faster
and avoid unnecessary memory allocations.

Do note that if you are interested in the indices, ranging over a
string and over a slice of runes will yield different indices. The
first one yields byte offsets, while the second one yields indices in
the slice of runes.`,
		Since:    "2017.1",
		Severity: lint.SeverityWarning,
		MergeIf:  lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

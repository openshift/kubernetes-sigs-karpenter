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

package analysisutil

import (
	"go/token"

	"github.com/gostaticanalysis/comment"
	"github.com/gostaticanalysis/comment/passes/commentmap"
	"golang.org/x/tools/go/analysis"
)

// ReportWithoutIgnore returns a report function which can set to (analysis.Pass).Report.
// The report function ignores a diagnostic which annotated by ignore comment as the below.
//   //lint:ignore Check1[,Check2,...,CheckN] reason
// names is a list of checker names.
// If names was omitted, the report function ignores by pass.Analyzer.Name.
func ReportWithoutIgnore(pass *analysis.Pass, names ...string) func(analysis.Diagnostic) {
	cmaps, _ := pass.ResultOf[commentmap.Analyzer].(comment.Maps)
	if cmaps == nil {
		cmaps = comment.New(pass.Fset, pass.Files)
	}

	if len(names) == 0 {
		names = []string{pass.Analyzer.Name}
	}

	report := pass.Report // original report func

	return func(d analysis.Diagnostic) {
		start := pass.Fset.File(d.Pos).Line(d.Pos)
		end := start
		if d.End != token.NoPos {
			end = pass.Fset.File(d.End).Line(d.End)
		}

		for l := start; l <= end; l++ {
			for _, n := range names {
				if cmaps.IgnoreLine(pass.Fset, l, n) {
					return
				}
			}
		}

		report(d)
	}
}

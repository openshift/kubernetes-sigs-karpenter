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

package prealloc

import (
	"fmt"

	"github.com/alexkohler/prealloc/pkg"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
)

func New(settings *config.PreallocSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name: "prealloc",
			Doc:  "Find slice declarations that could potentially be pre-allocated",
			Run: func(pass *analysis.Pass) (any, error) {
				runPreAlloc(pass, settings)

				return nil, nil
			},
		}).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

func runPreAlloc(pass *analysis.Pass, settings *config.PreallocSettings) {
	hints := pkg.Check(pass.Files, settings.Simple, settings.RangeLoops, settings.ForLoops)

	for _, hint := range hints {
		pass.Report(analysis.Diagnostic{
			Pos:     hint.Pos,
			Message: fmt.Sprintf("Consider pre-allocating %s", internal.FormatCode(hint.DeclaredSliceName)),
		})
	}
}

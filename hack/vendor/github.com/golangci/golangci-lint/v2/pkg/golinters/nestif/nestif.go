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

package nestif

import (
	"github.com/nakabonne/nestif"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.NestifSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name: "nestif",
			Doc:  "Reports deeply nested if statements",
			Run: func(pass *analysis.Pass) (any, error) {
				runNestIf(pass, settings)

				return nil, nil
			},
		}).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

func runNestIf(pass *analysis.Pass, settings *config.NestifSettings) {
	checker := &nestif.Checker{
		MinComplexity: settings.MinComplexity,
	}

	for _, file := range pass.Files {
		position, isGoFile := goanalysis.GetGoFilePosition(pass, file)
		if !isGoFile {
			continue
		}

		issues := checker.Check(file, pass.Fset)
		if len(issues) == 0 {
			continue
		}

		nonAdjPosition := pass.Fset.PositionFor(file.Pos(), false)

		f := pass.Fset.File(file.Pos())

		for _, issue := range issues {
			pass.Report(analysis.Diagnostic{
				Pos:     f.LineStart(goanalysis.AdjustPos(issue.Pos.Line, nonAdjPosition.Line, position.Line)),
				Message: issue.Message,
			})
		}
	}
}

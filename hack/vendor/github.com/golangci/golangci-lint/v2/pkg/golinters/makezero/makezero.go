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

package makezero

import (
	"fmt"

	"github.com/ashanbrown/makezero/v2/makezero"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.MakezeroSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name: "makezero",
			Doc:  "Find slice declarations with non-zero initial length",
			Run: func(pass *analysis.Pass) (any, error) {
				err := runMakeZero(pass, settings)
				if err != nil {
					return nil, err
				}

				return nil, nil
			},
		}).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

func runMakeZero(pass *analysis.Pass, settings *config.MakezeroSettings) error {
	zero := makezero.NewLinter(settings.Always)

	for _, file := range pass.Files {
		hints, err := zero.Run(pass.Fset, pass.TypesInfo, file)
		if err != nil {
			return fmt.Errorf("makezero linter failed on file %q: %w", file.Name.String(), err)
		}

		for _, hint := range hints {
			pass.Report(analysis.Diagnostic{
				Pos:     hint.Pos(),
				Message: hint.Details(),
			})
		}
	}

	return nil
}

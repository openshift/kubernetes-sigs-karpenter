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

package godot

import (
	"cmp"

	"github.com/tetafro/godot"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.GodotSettings) *goanalysis.Linter {
	var dotSettings godot.Settings

	if settings != nil {
		dotSettings = godot.Settings{
			Scope:   godot.Scope(settings.Scope),
			Exclude: settings.Exclude,
			Period:  settings.Period,
			Capital: settings.Capital,
		}
	}

	dotSettings.Scope = cmp.Or(dotSettings.Scope, godot.DeclScope)

	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name: "godot",
			Doc:  "Check if comments end in a period",
			Run: func(pass *analysis.Pass) (any, error) {
				err := runGodot(pass, dotSettings)
				if err != nil {
					return nil, err
				}

				return nil, nil
			},
		}).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

func runGodot(pass *analysis.Pass, settings godot.Settings) error {
	for _, file := range pass.Files {
		iss, err := godot.Run(file, pass.Fset, settings)
		if err != nil {
			return err
		}

		if len(iss) == 0 {
			continue
		}

		f := pass.Fset.File(file.Pos())

		for _, i := range iss {
			start := f.Pos(i.Pos.Offset)
			end := goanalysis.EndOfLinePos(f, i.Pos.Line)

			pass.Report(analysis.Diagnostic{
				Pos:     start,
				End:     end,
				Message: i.Message,
				SuggestedFixes: []analysis.SuggestedFix{{
					TextEdits: []analysis.TextEdit{{
						Pos:     start,
						End:     end,
						NewText: []byte(i.Replacement),
					}},
				}},
			})
		}
	}

	return nil
}

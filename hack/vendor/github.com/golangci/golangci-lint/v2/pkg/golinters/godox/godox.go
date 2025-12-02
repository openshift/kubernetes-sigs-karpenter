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

package godox

import (
	"go/token"
	"strings"

	"github.com/matoous/godox"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.GodoxSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name: "godox",
			Doc:  "Detects usage of FIXME, TODO and other keywords inside comments",
			Run: func(pass *analysis.Pass) (any, error) {
				return run(pass, settings), nil
			},
		}).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

func run(pass *analysis.Pass, settings *config.GodoxSettings) error {
	for _, file := range pass.Files {
		position, isGoFile := goanalysis.GetGoFilePosition(pass, file)
		if !isGoFile {
			continue
		}

		messages, err := godox.Run(file, pass.Fset, settings.Keywords...)
		if err != nil {
			return err
		}

		if len(messages) == 0 {
			continue
		}

		nonAdjPosition := pass.Fset.PositionFor(file.Pos(), false)

		ft := pass.Fset.File(file.Pos())

		for _, i := range messages {
			msg := strings.TrimRight(i.Message, "\n")

			// https://github.com/matoous/godox/blob/1d6ac9d899726279072e1c4d2b10f1eb52f22878/godox.go#L56
			index := strings.Index(msg, "Line contains")

			pass.Report(analysis.Diagnostic{
				Pos:     ft.LineStart(goanalysis.AdjustPos(i.Pos.Line, nonAdjPosition.Line, position.Line)) + token.Pos(i.Pos.Column),
				Message: msg[index:],
			})
		}
	}

	return nil
}

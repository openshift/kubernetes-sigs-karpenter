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

package nolintlint

import (
	"fmt"
	"sync"

	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
	nolintlint "github.com/golangci/golangci-lint/v2/pkg/golinters/nolintlint/internal"
	"github.com/golangci/golangci-lint/v2/pkg/lint/linter"
)

const LinterName = nolintlint.LinterName

func New(settings *config.NoLintLintSettings) *goanalysis.Linter {
	var mu sync.Mutex
	var resIssues []*goanalysis.Issue

	var needs nolintlint.Needs
	if settings.RequireExplanation {
		needs |= nolintlint.NeedsExplanation
	}
	if settings.RequireSpecific {
		needs |= nolintlint.NeedsSpecific
	}
	if !settings.AllowUnused {
		needs |= nolintlint.NeedsUnused
	}

	lnt, err := nolintlint.NewLinter(needs, settings.AllowNoExplanation)
	if err != nil {
		internal.LinterLogger.Fatalf("%s: create analyzer: %v", nolintlint.LinterName, err)
	}

	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name: nolintlint.LinterName,
			Doc:  "Reports ill-formed or insufficient nolint directives",
			Run: func(pass *analysis.Pass) (any, error) {
				issues, err := lnt.Run(pass)
				if err != nil {
					return nil, fmt.Errorf("linter failed to run: %w", err)
				}

				if len(issues) == 0 {
					return nil, nil
				}

				mu.Lock()
				resIssues = append(resIssues, issues...)
				mu.Unlock()

				return nil, nil
			},
		}).
		WithIssuesReporter(func(*linter.Context) []*goanalysis.Issue {
			return resIssues
		}).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

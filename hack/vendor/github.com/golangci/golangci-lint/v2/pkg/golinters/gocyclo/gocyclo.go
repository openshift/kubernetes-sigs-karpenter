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

package gocyclo

import (
	"fmt"
	"sync"

	"github.com/fzipp/gocyclo"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
	"github.com/golangci/golangci-lint/v2/pkg/lint/linter"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

const linterName = "gocyclo"

func New(settings *config.GoCycloSettings) *goanalysis.Linter {
	var mu sync.Mutex
	var resIssues []*goanalysis.Issue

	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name: linterName,
			Doc:  "Computes and checks the cyclomatic complexity of functions",
			Run: func(pass *analysis.Pass) (any, error) {
				issues := runGoCyclo(pass, settings)

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

func runGoCyclo(pass *analysis.Pass, settings *config.GoCycloSettings) []*goanalysis.Issue {
	var stats gocyclo.Stats
	for _, f := range pass.Files {
		stats = gocyclo.AnalyzeASTFile(f, pass.Fset, stats)
	}
	if len(stats) == 0 {
		return nil
	}

	stats = stats.SortAndFilter(-1, settings.MinComplexity)

	issues := make([]*goanalysis.Issue, 0, len(stats))

	for _, s := range stats {
		text := fmt.Sprintf("cyclomatic complexity %d of func %s is high (> %d)",
			s.Complexity, internal.FormatCode(s.FuncName), settings.MinComplexity)

		issues = append(issues, goanalysis.NewIssue(&result.Issue{
			Pos:        s.Pos,
			Text:       text,
			FromLinter: linterName,
		}, pass))
	}

	return issues
}

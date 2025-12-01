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

package promlinter

import (
	"fmt"
	"sync"

	"github.com/yeya24/promlinter"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/lint/linter"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

const linterName = "promlinter"

func New(settings *config.PromlinterSettings) *goanalysis.Linter {
	var mu sync.Mutex
	var resIssues []*goanalysis.Issue

	var promSettings promlinter.Setting
	if settings != nil {
		promSettings = promlinter.Setting{
			Strict:            settings.Strict,
			DisabledLintFuncs: settings.DisabledLinters,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name: linterName,
			Doc:  "Check Prometheus metrics naming via promlint",
			Run: func(pass *analysis.Pass) (any, error) {
				issues := runPromLinter(pass, promSettings)

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

func runPromLinter(pass *analysis.Pass, promSettings promlinter.Setting) []*goanalysis.Issue {
	lintIssues := promlinter.RunLint(pass.Fset, pass.Files, promSettings)

	if len(lintIssues) == 0 {
		return nil
	}

	issues := make([]*goanalysis.Issue, len(lintIssues))
	for k, i := range lintIssues {
		issue := result.Issue{
			Pos:        i.Pos,
			Text:       fmt.Sprintf("Metric: %s Error: %s", i.Metric, i.Text),
			FromLinter: linterName,
		}

		issues[k] = goanalysis.NewIssue(&issue, pass)
	}

	return issues
}

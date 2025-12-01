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

package goconst

import (
	"fmt"
	"sync"

	goconstAPI "github.com/jgautheron/goconst"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
	"github.com/golangci/golangci-lint/v2/pkg/lint/linter"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

const linterName = "goconst"

func New(settings *config.GoConstSettings) *goanalysis.Linter {
	var mu sync.Mutex
	var resIssues []*goanalysis.Issue

	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name: linterName,
			Doc:  "Finds repeated strings that could be replaced by a constant",
			Run: func(pass *analysis.Pass) (any, error) {
				issues, err := runGoconst(pass, settings)
				if err != nil {
					return nil, err
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
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

func runGoconst(pass *analysis.Pass, settings *config.GoConstSettings) ([]*goanalysis.Issue, error) {
	cfg := goconstAPI.Config{
		IgnoreStrings:        settings.IgnoreStringValues,
		MatchWithConstants:   settings.MatchWithConstants,
		MinStringLength:      settings.MinStringLen,
		MinOccurrences:       settings.MinOccurrencesCount,
		ParseNumbers:         settings.ParseNumbers,
		NumberMin:            settings.NumberMin,
		NumberMax:            settings.NumberMax,
		ExcludeTypes:         map[goconstAPI.Type]bool{},
		FindDuplicates:       settings.FindDuplicates,
		EvalConstExpressions: settings.EvalConstExpressions,

		// Should be managed with `linters.exclusions.rules`.
		IgnoreTests: false,
	}

	if settings.IgnoreCalls {
		cfg.ExcludeTypes[goconstAPI.Call] = true
	}

	lintIssues, err := goconstAPI.Run(pass.Files, pass.Fset, pass.TypesInfo, &cfg)
	if err != nil {
		return nil, err
	}

	if len(lintIssues) == 0 {
		return nil, nil
	}

	res := make([]*goanalysis.Issue, 0, len(lintIssues))
	for i := range lintIssues {
		issue := &lintIssues[i]

		var text string

		switch {
		case issue.OccurrencesCount > 0:
			text = fmt.Sprintf("string %s has %d occurrences", internal.FormatCode(issue.Str), issue.OccurrencesCount)

			if issue.MatchingConst == "" {
				text += ", make it a constant"
			} else {
				text += fmt.Sprintf(", but such constant %s already exists", internal.FormatCode(issue.MatchingConst))
			}

		case issue.DuplicateConst != "":
			text = fmt.Sprintf("This constant is a duplicate of %s at %s",
				internal.FormatCode(issue.DuplicateConst),
				issue.DuplicatePos.String())

		default:
			continue
		}

		res = append(res, goanalysis.NewIssue(&result.Issue{
			Pos:        issue.Pos,
			Text:       text,
			FromLinter: linterName,
		}, pass))
	}

	return res, nil
}

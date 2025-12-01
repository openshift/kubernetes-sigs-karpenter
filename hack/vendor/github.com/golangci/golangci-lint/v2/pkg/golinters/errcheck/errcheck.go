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

package errcheck

import (
	"cmp"
	"fmt"
	"regexp"
	"sync"

	"github.com/kisielk/errcheck/errcheck"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/packages"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
	"github.com/golangci/golangci-lint/v2/pkg/lint/linter"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

const linterName = "errcheck"

func New(settings *config.ErrcheckSettings) *goanalysis.Linter {
	var mu sync.Mutex
	var resIssues []*goanalysis.Issue

	analyzer := &analysis.Analyzer{
		Name: linterName,
		Doc: "errcheck is a program for checking for unchecked errors in Go code. " +
			"These unchecked errors can be critical bugs in some cases",
		Run: goanalysis.DummyRun,
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer).
		WithContextSetter(func(lintCtx *linter.Context) {
			checker := getChecker(settings)
			checker.Tags = lintCtx.Cfg.Run.BuildTags

			analyzer.Run = func(pass *analysis.Pass) (any, error) {
				issues := runErrCheck(pass, checker, settings.Verbose)

				if len(issues) == 0 {
					return nil, nil
				}

				mu.Lock()
				resIssues = append(resIssues, issues...)
				mu.Unlock()

				return nil, nil
			}
		}).
		WithIssuesReporter(func(*linter.Context) []*goanalysis.Issue {
			return resIssues
		}).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

func runErrCheck(pass *analysis.Pass, checker *errcheck.Checker, verbose bool) []*goanalysis.Issue {
	pkg := &packages.Package{
		Fset:      pass.Fset,
		Syntax:    pass.Files,
		Types:     pass.Pkg,
		TypesInfo: pass.TypesInfo,
	}

	lintIssues := checker.CheckPackage(pkg).Unique()
	if len(lintIssues.UncheckedErrors) == 0 {
		return nil
	}

	issues := make([]*goanalysis.Issue, len(lintIssues.UncheckedErrors))

	for i, err := range lintIssues.UncheckedErrors {
		text := "Error return value is not checked"

		if err.FuncName != "" {
			code := cmp.Or(err.SelectorName, err.FuncName)
			if verbose {
				code = err.FuncName
			}

			text = fmt.Sprintf("Error return value of %s is not checked", internal.FormatCode(code))
		}

		issues[i] = goanalysis.NewIssue(
			&result.Issue{
				FromLinter: linterName,
				Text:       text,
				Pos:        err.Pos,
			},
			pass,
		)
	}

	return issues
}

func getChecker(errCfg *config.ErrcheckSettings) *errcheck.Checker {
	checker := errcheck.Checker{
		Exclusions: errcheck.Exclusions{
			BlankAssignments:       !errCfg.CheckAssignToBlank,
			TypeAssertions:         !errCfg.CheckTypeAssertions,
			SymbolRegexpsByPackage: map[string]*regexp.Regexp{},
		},
	}

	if !errCfg.DisableDefaultExclusions {
		checker.Exclusions.Symbols = append(checker.Exclusions.Symbols, errcheck.DefaultExcludedSymbols...)
	}

	checker.Exclusions.Symbols = append(checker.Exclusions.Symbols, errCfg.ExcludeFunctions...)

	return &checker
}

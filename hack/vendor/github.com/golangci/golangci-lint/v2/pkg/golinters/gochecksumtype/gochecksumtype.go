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

package gochecksumtype

import (
	"strings"
	"sync"

	gochecksumtype "github.com/alecthomas/go-check-sumtype"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/packages"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/lint/linter"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

const linterName = "gochecksumtype"

func New(settings *config.GoChecksumTypeSettings) *goanalysis.Linter {
	var mu sync.Mutex
	var resIssues []*goanalysis.Issue

	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name: linterName,
			Doc:  `Run exhaustiveness checks on Go "sum types"`,
			Run: func(pass *analysis.Pass) (any, error) {
				issues, err := runGoCheckSumType(pass, settings)
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
		WithIssuesReporter(func(_ *linter.Context) []*goanalysis.Issue {
			return resIssues
		}).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

func runGoCheckSumType(pass *analysis.Pass, settings *config.GoChecksumTypeSettings) ([]*goanalysis.Issue, error) {
	var resIssues []*goanalysis.Issue

	pkg := &packages.Package{
		Fset:      pass.Fset,
		Syntax:    pass.Files,
		Types:     pass.Pkg,
		TypesInfo: pass.TypesInfo,
	}

	cfg := gochecksumtype.Config{
		DefaultSignifiesExhaustive: settings.DefaultSignifiesExhaustive,
		IncludeSharedInterfaces:    settings.IncludeSharedInterfaces,
	}

	var unknownError error
	errors := gochecksumtype.Run([]*packages.Package{pkg}, cfg)
	for _, err := range errors {
		err, ok := err.(gochecksumtype.Error)
		if !ok {
			unknownError = err
			continue
		}

		resIssues = append(resIssues, goanalysis.NewIssue(&result.Issue{
			FromLinter: linterName,
			Text:       strings.TrimPrefix(err.Error(), err.Pos().String()+": "),
			Pos:        err.Pos(),
		}, pass))
	}

	return resIssues, unknownError
}

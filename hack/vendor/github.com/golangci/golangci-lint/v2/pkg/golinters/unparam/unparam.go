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

package unparam

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/packages"
	"mvdan.cc/unparam/check"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.UnparamSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(&analysis.Analyzer{
			Name:     "unparam",
			Doc:      "Reports unused function parameters",
			Requires: []*analysis.Analyzer{buildssa.Analyzer},
			Run: func(pass *analysis.Pass) (any, error) {
				err := runUnparam(pass, settings)
				if err != nil {
					return nil, err
				}

				return nil, nil
			},
		}).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

func runUnparam(pass *analysis.Pass, settings *config.UnparamSettings) error {
	ssa := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	ssaPkg := ssa.Pkg

	pkg := &packages.Package{
		Fset:      pass.Fset,
		Syntax:    pass.Files,
		Types:     pass.Pkg,
		TypesInfo: pass.TypesInfo,
	}

	c := &check.Checker{}
	c.CheckExportedFuncs(settings.CheckExported)
	c.Packages([]*packages.Package{pkg})
	c.ProgramSSA(ssaPkg.Prog)

	unparamIssues, err := c.Check()
	if err != nil {
		return err
	}

	for _, i := range unparamIssues {
		pass.Report(analysis.Diagnostic{
			Pos:     i.Pos(),
			Message: i.Message(),
		})
	}

	return nil
}

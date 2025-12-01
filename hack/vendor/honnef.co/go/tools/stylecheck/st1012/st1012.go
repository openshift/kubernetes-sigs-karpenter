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

package st1012

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/analysis/report"

	"golang.org/x/tools/go/analysis"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name: "ST1012",
		Run:  run,
	},
	Doc: &lint.RawDocumentation{
		Title: `Poorly chosen name for error variable`,
		Text: `Error variables that are part of an API should be called \'errFoo\' or
\'ErrFoo\'.`,
		Since:   "2019.1",
		MergeIf: lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.VAR {
				continue
			}
			for _, spec := range gen.Specs {
				spec := spec.(*ast.ValueSpec)
				if len(spec.Names) != len(spec.Values) {
					continue
				}

				for i, name := range spec.Names {
					val := spec.Values[i]
					if !code.IsCallToAny(pass, val, "errors.New", "fmt.Errorf") {
						continue
					}

					if pass.Pkg.Path() == "net/http" && strings.HasPrefix(name.Name, "http2err") {
						// special case for internal variable names of
						// bundled HTTP 2 code in net/http
						continue
					}
					prefix := "err"
					if name.IsExported() {
						prefix = "Err"
					}
					if !strings.HasPrefix(name.Name, prefix) {
						report.Report(pass, name, fmt.Sprintf("error var %s should have name of the form %sFoo", name.Name, prefix))
					}
				}
			}
		}
	}
	return nil, nil
}

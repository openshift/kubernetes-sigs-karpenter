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

package testableexamples

import (
	"go/ast"
	"go/doc"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// NewAnalyzer returns Analyzer that checks if examples are testable.
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "testableexamples",
		Doc:  "linter checks if examples are testable (have an expected output)",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			testFiles := make([]*ast.File, 0, len(pass.Files))
			for _, file := range pass.Files {
				fileName := pass.Fset.File(file.Pos()).Name()
				if strings.HasSuffix(fileName, "_test.go") {
					testFiles = append(testFiles, file)
				}
			}

			for _, example := range doc.Examples(testFiles...) {
				if example.Output == "" && !example.EmptyOutput {
					pass.Reportf(example.Code.Pos(), "missing output for example, go test can't validate it")
				}
			}

			return nil, nil
		},
	}
}

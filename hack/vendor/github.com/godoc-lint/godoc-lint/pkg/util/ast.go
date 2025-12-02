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

package util

import (
	"go/ast"
	"go/token"
	"iter"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// GetPassFileToken is a helper function to return the file token associated
// with the given AST file.
func GetPassFileToken(f *ast.File, pass *analysis.Pass) *token.File {
	if f.Pos() == token.NoPos {
		return nil
	}
	ft := pass.Fset.File(f.Pos())
	if ft == nil {
		return nil
	}
	return ft
}

// AnalysisApplicableFiles returns an iterator looping over files that are ready
// to be analyzed.
//
// The yield-ed arguments are never nil.
func AnalysisApplicableFiles(actx *model.AnalysisContext, includeTests bool, ruleSet model.RuleSet) iter.Seq2[*ast.File, *model.FileInspection] {
	return func(yield func(*ast.File, *model.FileInspection) bool) {
		if actx.InspectorResult == nil {
			return
		}

		for _, f := range actx.Pass.Files {
			ir := actx.InspectorResult.Files[f]

			if ir == nil {
				continue
			}

			ft := GetPassFileToken(f, actx.Pass)
			if ft == nil {
				continue
			}

			if !actx.Config.IsPathApplicable(ft.Name()) {
				continue
			}

			if !includeTests && strings.HasSuffix(ft.Name(), "_test.go") {
				continue
			}

			if ir.DisabledRules.All || ir.DisabledRules.Rules.IsSupersetOf(ruleSet) {
				continue
			}

			if !yield(f, ir) {
				return
			}
		}
	}
}

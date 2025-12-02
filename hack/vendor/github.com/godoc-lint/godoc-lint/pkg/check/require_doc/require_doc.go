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

// Package require_doc provides a checker that requires symbols to have godocs.
package require_doc

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/godoc-lint/godoc-lint/pkg/model"
	"github.com/godoc-lint/godoc-lint/pkg/util"
)

const requireDocRule = model.RequireDocRule

var ruleSet = model.RuleSet{}.Add(requireDocRule)

// RequireDocChecker checks required godocs.
type RequireDocChecker struct{}

// NewRequireDocChecker returns a new instance of the corresponding checker.
func NewRequireDocChecker() *RequireDocChecker {
	return &RequireDocChecker{}
}

// GetCoveredRules implements the corresponding interface method.
func (r *RequireDocChecker) GetCoveredRules() model.RuleSet {
	return ruleSet
}

// Apply implements the corresponding interface method.
func (r *RequireDocChecker) Apply(actx *model.AnalysisContext) error {
	includeTests := actx.Config.GetRuleOptions().RequireDocIncludeTests
	requirePublic := !actx.Config.GetRuleOptions().RequireDocIgnoreExported
	requirePrivate := !actx.Config.GetRuleOptions().RequireDocIgnoreUnexported

	if !requirePublic && !requirePrivate {
		return nil
	}

	for _, ir := range util.AnalysisApplicableFiles(actx, includeTests, model.RuleSet{}.Add(requireDocRule)) {
		for _, decl := range ir.SymbolDecl {
			isExported := ast.IsExported(decl.Name)
			if isExported && !requirePublic || !isExported && !requirePrivate {
				continue
			}

			if decl.Doc != nil && (decl.Doc.DisabledRules.All || decl.Doc.DisabledRules.Rules.Has(requireDocRule)) {
				continue
			}

			if decl.Kind == model.SymbolDeclKindBad {
				continue
			}

			if decl.Kind == model.SymbolDeclKindFunc {
				if decl.Doc == nil || decl.Doc.Text == "" {
					reportRange(actx.Pass, decl.Ident)
				}
				continue
			}

			// Now we only have const/var/type declarations.

			if decl.Doc != nil && decl.Doc.Text != "" {
				// cases:
				//
				//   // godoc
				//   const foo = 0
				//
				//   // godoc
				//   const foo, bar = 0, 0
				//
				//   const (
				//       // godoc
				//       foo = 0
				//   )
				//
				//   const (
				//       // godoc
				//       foo, bar = 0, 0
				//   )
				//
				//   // godoc
				//   type foo int
				//
				//   type (
				//       // godoc
				//       foo int
				//   )
				continue
			}

			if decl.TrailingDoc != nil && decl.TrailingDoc.Text != "" {
				// cases:
				//
				//   const foo = 0 // godoc
				//
				//   const foo, bar = 0, 0 // godoc
				//
				//   const (
				//       foo = 0 // godoc
				//   )
				//
				//   const (
				//       foo, bar = 0, 0 // godoc
				//   )
				//
				//   type foo int // godoc
				//
				//   type (
				//       foo int // godoc
				//   )
				continue
			}

			if decl.ParentDoc != nil && decl.ParentDoc.Text != "" {
				// cases:
				//
				//   // godoc
				//   const (
				//       foo = 0
				//   )
				//
				//   // godoc
				//   const (
				//       foo, bar = 0, 0
				//   )
				//
				//   // godoc
				//   type (
				//       foo int
				//   )
				continue
			}

			// At this point there is no godoc for the symbol.
			//
			// cases:
			//
			//   const foo = 0
			//
			//   const foo, bar = 0, 0
			//
			//   const (
			//       foo = 0
			//   )
			//
			//   const (
			//       foo, bar = 0, 0
			//   )
			//
			//   type foo int
			//
			//   type (
			//       foo int
			//   )

			reportRange(actx.Pass, decl.Ident)
		}
	}
	return nil
}

func reportRange(pass *analysis.Pass, ident *ast.Ident) {
	pass.ReportRangef(ident, "symbol should have a godoc (%q)", ident.Name)
}

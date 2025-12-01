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

// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modernize

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/internal/analysisinternal"
	"golang.org/x/tools/internal/analysisinternal/generated"
)

var AnyAnalyzer = &analysis.Analyzer{
	Name: "any",
	Doc:  analysisinternal.MustExtractDoc(doc, "any"),
	Requires: []*analysis.Analyzer{
		generated.Analyzer,
		inspect.Analyzer,
	},
	Run: runAny,
	URL: "https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/modernize#any",
}

// The any pass replaces interface{} with go1.18's 'any'.
func runAny(pass *analysis.Pass) (any, error) {
	skipGenerated(pass)

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	for curFile := range filesUsing(inspect, pass.TypesInfo, "go1.18") {
		for curIface := range curFile.Preorder((*ast.InterfaceType)(nil)) {
			iface := curIface.Node().(*ast.InterfaceType)

			if iface.Methods.NumFields() == 0 {
				// Check that 'any' is not shadowed.
				if lookup(pass.TypesInfo, curIface, "any") == builtinAny {
					pass.Report(analysis.Diagnostic{
						Pos:     iface.Pos(),
						End:     iface.End(),
						Message: "interface{} can be replaced by any",
						SuggestedFixes: []analysis.SuggestedFix{{
							Message: "Replace interface{} by any",
							TextEdits: []analysis.TextEdit{
								{
									Pos:     iface.Pos(),
									End:     iface.End(),
									NewText: []byte("any"),
								},
							},
						}},
					})
				}
			}
		}
	}
	return nil, nil
}

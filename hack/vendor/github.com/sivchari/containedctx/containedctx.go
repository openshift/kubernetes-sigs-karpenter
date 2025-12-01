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

package containedctx

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "containedctx is a linter that detects struct contained context.Context field"

// Analyzer is the contanedctx analyzer
var Analyzer = &analysis.Analyzer{
	Name: "containedctx",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch structTyp := n.(type) {
		case *ast.StructType:
			if structTyp.Fields.List == nil {
				return
			}
			for _, field := range structTyp.Fields.List {
				if pass.TypesInfo.TypeOf(field.Type).String() == "context.Context" {
					pass.Reportf(field.Pos(), "found a struct that contains a context.Context field")
				}
			}
		}
	})

	return nil, nil
}

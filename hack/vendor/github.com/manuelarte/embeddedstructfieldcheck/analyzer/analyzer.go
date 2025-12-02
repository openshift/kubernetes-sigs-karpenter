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

package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/manuelarte/embeddedstructfieldcheck/internal"
)

const (
	EmptyLineCheck   = "empty-line"
	ForbidMutexCheck = "forbid-mutex"
)

func NewAnalyzer() *analysis.Analyzer {
	var (
		emptyLine   bool
		forbidMutex bool
	)

	a := &analysis.Analyzer{
		Name: "embeddedstructfieldcheck",
		Doc: "Embedded types should be at the top of the field list of a struct, " +
			"and there must be an empty line separating embedded fields from regular fields.",
		URL: "https://github.com/manuelarte/embeddedstructfieldcheck",
		Run: func(pass *analysis.Pass) (any, error) {
			run(pass, emptyLine, forbidMutex)

			//nolint:nilnil // impossible case.
			return nil, nil
		},
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	a.Flags.BoolVar(&emptyLine, EmptyLineCheck, true,
		"Checks that there is an empty space between the embedded fields and regular fields.")
	a.Flags.BoolVar(&forbidMutex, ForbidMutexCheck, false,
		"Checks that sync.Mutex and sync.RWMutex are not used as an embedded fields.")

	return a
}

func run(pass *analysis.Pass, emptyLine, forbidMutex bool) {
	insp, found := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !found {
		return
	}

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		node, ok := n.(*ast.StructType)
		if !ok {
			return
		}

		internal.Analyze(pass, node, emptyLine, forbidMutex)
	})
}

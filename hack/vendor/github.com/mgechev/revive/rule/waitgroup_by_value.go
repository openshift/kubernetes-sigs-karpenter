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

package rule

import (
	"go/ast"

	"github.com/mgechev/revive/internal/astutils"
	"github.com/mgechev/revive/lint"
)

// WaitGroupByValueRule lints sync.WaitGroup passed by copy in functions.
type WaitGroupByValueRule struct{}

// Apply applies the rule to given file.
func (*WaitGroupByValueRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}

	w := lintWaitGroupByValueRule{onFailure: onFailure}
	ast.Walk(w, file.AST)
	return failures
}

// Name returns the rule name.
func (*WaitGroupByValueRule) Name() string {
	return "waitgroup-by-value"
}

type lintWaitGroupByValueRule struct {
	onFailure func(lint.Failure)
}

func (w lintWaitGroupByValueRule) Visit(node ast.Node) ast.Visitor {
	// look for function declarations
	fd, ok := node.(*ast.FuncDecl)
	if !ok {
		return w
	}

	// Check all function parameters
	for _, field := range fd.Type.Params.List {
		if !astutils.IsPkgDotName(field.Type, "sync", "WaitGroup") {
			continue
		}

		w.onFailure(lint.Failure{
			Confidence: 1,
			Node:       field,
			Failure:    "sync.WaitGroup passed by value, the function will get a copy of the original one",
		})
	}

	return nil // skip visiting function body
}

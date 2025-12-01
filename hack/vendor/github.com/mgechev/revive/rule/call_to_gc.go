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

// CallToGCRule lints calls to the garbage collector.
type CallToGCRule struct{}

// Apply applies the rule to given file.
func (*CallToGCRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure
	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}

	w := lintCallToGC{onFailure}
	ast.Walk(w, file.AST)

	return failures
}

// Name returns the rule name.
func (*CallToGCRule) Name() string {
	return "call-to-gc"
}

type lintCallToGC struct {
	onFailure func(lint.Failure)
}

func (w lintCallToGC) Visit(node ast.Node) ast.Visitor {
	ce, ok := node.(*ast.CallExpr)
	if !ok {
		return w // nothing to do, the node is not a function call
	}

	if !astutils.IsPkgDotName(ce.Fun, "runtime", "GC") {
		return nil // nothing to do, the call is not a call to the Garbage Collector
	}

	w.onFailure(lint.Failure{
		Confidence: 1,
		Node:       node,
		Category:   lint.FailureCategoryBadPractice,
		Failure:    "explicit call to the garbage collector",
	})

	return w
}

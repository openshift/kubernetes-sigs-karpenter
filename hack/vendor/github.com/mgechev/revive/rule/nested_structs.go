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

	"github.com/mgechev/revive/lint"
)

// NestedStructs lints nested structs.
type NestedStructs struct{}

// Apply applies the rule to given file.
func (*NestedStructs) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	walker := &lintNestedStructs{
		onFailure: func(failure lint.Failure) {
			failures = append(failures, failure)
		},
	}

	ast.Walk(walker, file.AST)

	return failures
}

// Name returns the rule name.
func (*NestedStructs) Name() string {
	return "nested-structs"
}

type lintNestedStructs struct {
	onFailure func(lint.Failure)
}

func (l *lintNestedStructs) Visit(n ast.Node) ast.Visitor {
	if v, ok := n.(*ast.StructType); ok {
		ls := &lintStruct{l.onFailure}
		ast.Walk(ls, v.Fields)
	}

	return l
}

type lintStruct struct {
	onFailure func(lint.Failure)
}

func (l *lintStruct) Visit(n ast.Node) ast.Visitor {
	switch s := n.(type) {
	case *ast.StructType:
		l.fail(s)
		return nil
	case *ast.ArrayType:
		if _, ok := s.Elt.(*ast.StructType); ok {
			l.fail(s)
		}
		return nil
	case *ast.ChanType:
		return nil
	case *ast.MapType:
		return nil
	default:
		return l
	}
}

func (l *lintStruct) fail(n ast.Node) {
	l.onFailure(lint.Failure{
		Failure:    "no nested structs are allowed",
		Category:   lint.FailureCategoryStyle,
		Node:       n,
		Confidence: 1,
	})
}

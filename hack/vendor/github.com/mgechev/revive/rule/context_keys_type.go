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
	"fmt"
	"go/ast"
	"go/types"

	"github.com/mgechev/revive/internal/astutils"
	"github.com/mgechev/revive/lint"
)

// ContextKeysType disallows the usage of basic types in `context.WithValue`.
type ContextKeysType struct{}

// Apply applies the rule to given file.
func (*ContextKeysType) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	fileAst := file.AST
	walker := lintContextKeyTypes{
		file:    file,
		fileAst: fileAst,
		onFailure: func(failure lint.Failure) {
			failures = append(failures, failure)
		},
	}

	file.Pkg.TypeCheck()
	ast.Walk(walker, fileAst)

	return failures
}

// Name returns the rule name.
func (*ContextKeysType) Name() string {
	return "context-keys-type"
}

type lintContextKeyTypes struct {
	file      *lint.File
	fileAst   *ast.File
	onFailure func(lint.Failure)
}

func (w lintContextKeyTypes) Visit(n ast.Node) ast.Visitor {
	if n, ok := n.(*ast.CallExpr); ok {
		checkContextKeyType(w, n)
	}

	return w
}

func checkContextKeyType(w lintContextKeyTypes, x *ast.CallExpr) {
	f := w.file
	if !astutils.IsPkgDotName(x.Fun, "context", "WithValue") {
		return
	}

	// key is second argument to context.WithValue
	if len(x.Args) != 3 {
		return
	}
	key := f.Pkg.TypesInfo().Types[x.Args[1]]

	if ktyp, ok := key.Type.(*types.Basic); ok && ktyp.Kind() != types.Invalid {
		w.onFailure(lint.Failure{
			Confidence: 1,
			Node:       x,
			Category:   lint.FailureCategoryContent,
			Failure:    fmt.Sprintf("should not use basic type %s as key in context.WithValue", key.Type),
		})
	}
}

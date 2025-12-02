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
	"strings"

	"github.com/mgechev/revive/lint"
)

// RedundantImportAlias warns on import aliases matching the imported package name.
type RedundantImportAlias struct{}

// Apply applies the rule to given file.
func (*RedundantImportAlias) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	for _, imp := range file.AST.Imports {
		if imp.Name == nil {
			continue
		}

		if getImportPackageName(imp) == imp.Name.Name {
			failures = append(failures, lint.Failure{
				Confidence: 1,
				Failure:    fmt.Sprintf("Import alias %q is redundant", imp.Name.Name),
				Node:       imp,
				Category:   lint.FailureCategoryImports,
			})
		}
	}

	return failures
}

// Name returns the rule name.
func (*RedundantImportAlias) Name() string {
	return "redundant-import-alias"
}

func getImportPackageName(imp *ast.ImportSpec) string {
	const pathSep = "/"
	const strDelim = `"`

	path := imp.Path.Value
	i := strings.LastIndex(path, pathSep)
	if i == -1 {
		return strings.Trim(path, strDelim)
	}

	return strings.Trim(path[i+1:], strDelim)
}

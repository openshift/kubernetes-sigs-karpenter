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

	"github.com/mgechev/revive/lint"
)

// DuplicatedImportsRule looks for packages that are imported two or more times.
type DuplicatedImportsRule struct{}

// Apply applies the rule to given file.
func (*DuplicatedImportsRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	impPaths := map[string]struct{}{}
	for _, imp := range file.AST.Imports {
		path := imp.Path.Value
		_, ok := impPaths[path]
		if ok {
			failures = append(failures, lint.Failure{
				Confidence: 1,
				Failure:    fmt.Sprintf("Package %s already imported", path),
				Node:       imp,
				Category:   lint.FailureCategoryImports,
			})
			continue
		}

		impPaths[path] = struct{}{}
	}

	return failures
}

// Name returns the rule name.
func (*DuplicatedImportsRule) Name() string {
	return "duplicated-imports"
}

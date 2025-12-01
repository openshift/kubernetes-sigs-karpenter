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

package mirror

import (
	"github.com/butuzov/mirror"

	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	// mirror only lints test files if the `--with-tests` flag is passed,
	// so we pass the `with-tests` flag as true to the analyzer before running it.
	// This can be turned off by using the regular golangci-lint flags such as `--tests` or `--skip-files`
	// or can be disabled per linter via exclude rules.
	// (see https://github.com/golangci/golangci-lint/issues/2527#issuecomment-1023707262)
	cfg := map[string]any{
		"with-tests": true,
	}

	return goanalysis.
		NewLinterFromAnalyzer(mirror.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

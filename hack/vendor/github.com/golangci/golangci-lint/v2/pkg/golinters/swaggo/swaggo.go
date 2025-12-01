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

package swaggo

import (
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters/swaggo"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(swaggo.Name),
				"Check if swaggo comments are formatted",
				swaggo.New(),
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

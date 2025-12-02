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

package modernize

import (
	"slices"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/modernize"
)

func New(settings *config.ModernizeSettings) *goanalysis.Linter {
	var analyzers []*analysis.Analyzer

	if settings == nil {
		analyzers = modernize.Suite
	} else {
		for _, analyzer := range modernize.Suite {
			if slices.Contains(settings.Disable, analyzer.Name) {
				continue
			}

			analyzers = append(analyzers, analyzer)
		}
	}

	return goanalysis.NewLinter(
		"modernize",
		"A suite of analyzers that suggest simplifications to Go code, using modern language and library features.",
		analyzers,
		nil).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

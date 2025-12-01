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

package mnd

import (
	mnd "github.com/tommy-muehle/go-mnd/v2"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.MndSettings) *goanalysis.Linter {
	cfg := map[string]any{}

	if settings != nil {
		if len(settings.Checks) > 0 {
			cfg["checks"] = settings.Checks
		}
		if len(settings.IgnoredNumbers) > 0 {
			cfg["ignored-numbers"] = settings.IgnoredNumbers
		}
		if len(settings.IgnoredFiles) > 0 {
			cfg["ignored-files"] = settings.IgnoredFiles
		}
		if len(settings.IgnoredFunctions) > 0 {
			cfg["ignored-functions"] = settings.IgnoredFunctions
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(mnd.Analyzer).
		WithDesc("An analyzer to detect magic numbers.").
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

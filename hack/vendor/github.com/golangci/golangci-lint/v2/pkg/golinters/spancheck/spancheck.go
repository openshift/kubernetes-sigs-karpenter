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

package spancheck

import (
	"github.com/jjti/go-spancheck"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.SpancheckSettings) *goanalysis.Linter {
	cfg := spancheck.NewDefaultConfig()

	if settings != nil {
		if len(settings.Checks) > 0 {
			cfg.EnabledChecks = settings.Checks
		}

		if len(settings.IgnoreCheckSignatures) > 0 {
			cfg.IgnoreChecksSignaturesSlice = settings.IgnoreCheckSignatures
		}

		if len(settings.ExtraStartSpanSignatures) > 0 {
			cfg.StartSpanMatchersSlice = append(cfg.StartSpanMatchersSlice, settings.ExtraStartSpanSignatures...)
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(spancheck.NewAnalyzerWithConfig(cfg)).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

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

package wrapcheck

import (
	"github.com/tomarrell/wrapcheck/v2/wrapcheck"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.WrapcheckSettings) *goanalysis.Linter {
	cfg := wrapcheck.NewDefaultConfig()

	if settings != nil {
		cfg.ExtraIgnoreSigs = settings.ExtraIgnoreSigs
		cfg.ReportInternalErrors = settings.ReportInternalErrors

		if len(settings.IgnoreSigs) != 0 {
			cfg.IgnoreSigs = settings.IgnoreSigs
		}
		if len(settings.IgnoreSigRegexps) != 0 {
			cfg.IgnoreSigRegexps = settings.IgnoreSigRegexps
		}
		if len(settings.IgnorePackageGlobs) != 0 {
			cfg.IgnorePackageGlobs = settings.IgnorePackageGlobs
		}
		if len(settings.IgnoreInterfaceRegexps) != 0 {
			cfg.IgnoreInterfaceRegexps = settings.IgnoreInterfaceRegexps
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(wrapcheck.NewAnalyzer(cfg)).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

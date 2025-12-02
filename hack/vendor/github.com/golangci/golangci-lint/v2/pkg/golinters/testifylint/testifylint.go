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

package testifylint

import (
	"github.com/Antonboom/testifylint/analyzer"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.TestifylintSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			"enable-all":  settings.EnableAll,
			"disable-all": settings.DisableAll,

			"bool-compare.ignore-custom-types": settings.BoolCompare.IgnoreCustomTypes,
			"formatter.require-f-funcs":        settings.Formatter.RequireFFuncs,
			"formatter.require-string-msg":     settings.Formatter.RequireStringMsg,
			"go-require.ignore-http-handlers":  settings.GoRequire.IgnoreHTTPHandlers,
		}
		if len(settings.EnabledCheckers) > 0 {
			cfg["enable"] = settings.EnabledCheckers
		}
		if len(settings.DisabledCheckers) > 0 {
			cfg["disable"] = settings.DisabledCheckers
		}

		if b := settings.Formatter.CheckFormatString; b != nil {
			cfg["formatter.check-format-string"] = *b
		}
		if p := settings.ExpectedActual.ExpVarPattern; p != "" {
			cfg["expected-actual.pattern"] = p
		}
		if p := settings.RequireError.FnPattern; p != "" {
			cfg["require-error.fn-pattern"] = p
		}
		if m := settings.SuiteExtraAssertCall.Mode; m != "" {
			cfg["suite-extra-assert-call.mode"] = m
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer.New()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

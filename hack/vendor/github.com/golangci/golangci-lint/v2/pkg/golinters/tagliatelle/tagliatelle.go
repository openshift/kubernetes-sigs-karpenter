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

package tagliatelle

import (
	"maps"
	"strings"

	"github.com/ldez/tagliatelle"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.TagliatelleSettings) *goanalysis.Linter {
	cfg := tagliatelle.Config{
		Base: tagliatelle.Base{
			Rules: map[string]string{
				"json":   "camel",
				"yaml":   "camel",
				"header": "header",
			},
		},
	}

	if settings != nil {
		maps.Copy(cfg.Rules, settings.Case.Rules)

		cfg.ExtendedRules = toExtendedRules(settings.Case.ExtendedRules)
		cfg.UseFieldName = settings.Case.UseFieldName
		cfg.IgnoredFields = settings.Case.IgnoredFields

		for _, override := range settings.Case.Overrides {
			cfg.Overrides = append(cfg.Overrides, tagliatelle.Overrides{
				Base: tagliatelle.Base{
					Rules:         override.Rules,
					ExtendedRules: toExtendedRules(override.ExtendedRules),
					UseFieldName:  override.UseFieldName,
					IgnoredFields: override.IgnoredFields,
					Ignore:        override.Ignore,
				},
				Package: override.Package,
			})
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(tagliatelle.New(cfg)).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

func toExtendedRules(src map[string]config.TagliatelleExtendedRule) map[string]tagliatelle.ExtendedRule {
	result := make(map[string]tagliatelle.ExtendedRule, len(src))

	for k, v := range src {
		initialismOverrides := make(map[string]bool, len(v.InitialismOverrides))
		for ki, vi := range v.InitialismOverrides {
			initialismOverrides[strings.ToUpper(ki)] = vi
		}

		result[k] = tagliatelle.ExtendedRule{
			Case:                v.Case,
			ExtraInitialisms:    v.ExtraInitialisms,
			InitialismOverrides: initialismOverrides,
		}
	}

	return result
}

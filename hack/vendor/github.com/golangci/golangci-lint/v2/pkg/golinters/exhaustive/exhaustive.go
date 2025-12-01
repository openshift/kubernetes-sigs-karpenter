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

package exhaustive

import (
	"github.com/nishanths/exhaustive"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.ExhaustiveSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			exhaustive.CheckFlag:                      settings.Check,
			exhaustive.DefaultSignifiesExhaustiveFlag: settings.DefaultSignifiesExhaustive,
			exhaustive.IgnoreEnumMembersFlag:          settings.IgnoreEnumMembers,
			exhaustive.IgnoreEnumTypesFlag:            settings.IgnoreEnumTypes,
			exhaustive.PackageScopeOnlyFlag:           settings.PackageScopeOnly,
			exhaustive.ExplicitExhaustiveMapFlag:      settings.ExplicitExhaustiveMap,
			exhaustive.ExplicitExhaustiveSwitchFlag:   settings.ExplicitExhaustiveSwitch,
			exhaustive.DefaultCaseRequiredFlag:        settings.DefaultCaseRequired,
			// Should be managed with `linters.exclusions.generated`.
			exhaustive.CheckGeneratedFlag: true,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(exhaustive.Analyzer).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

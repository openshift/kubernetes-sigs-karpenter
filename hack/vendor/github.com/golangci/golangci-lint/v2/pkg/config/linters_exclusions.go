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

package config

import (
	"fmt"
	"slices"
)

const (
	GeneratedModeLax     = "lax"
	GeneratedModeStrict  = "strict"
	GeneratedModeDisable = "disable"
)

const (
	ExclusionPresetComments             = "comments"
	ExclusionPresetStdErrorHandling     = "std-error-handling"
	ExclusionPresetCommonFalsePositives = "common-false-positives"
	ExclusionPresetLegacy               = "legacy"
)

const excludeRuleMinConditionsCount = 2

type LinterExclusions struct {
	Generated   string        `mapstructure:"generated"`
	WarnUnused  bool          `mapstructure:"warn-unused"`
	Presets     []string      `mapstructure:"presets"`
	Rules       []ExcludeRule `mapstructure:"rules"`
	Paths       []string      `mapstructure:"paths"`
	PathsExcept []string      `mapstructure:"paths-except"`
}

func (e *LinterExclusions) Validate() error {
	for i, rule := range e.Rules {
		if err := rule.Validate(); err != nil {
			return fmt.Errorf("error in exclude rule #%d: %w", i, err)
		}
	}

	allPresets := []string{
		ExclusionPresetComments,
		ExclusionPresetStdErrorHandling,
		ExclusionPresetCommonFalsePositives,
		ExclusionPresetLegacy,
	}

	for _, preset := range e.Presets {
		if !slices.Contains(allPresets, preset) {
			return fmt.Errorf("invalid preset: %s", preset)
		}
	}

	return nil
}

type ExcludeRule struct {
	BaseRule `mapstructure:",squash"`
}

func (e *ExcludeRule) Validate() error {
	return e.BaseRule.Validate(excludeRuleMinConditionsCount)
}

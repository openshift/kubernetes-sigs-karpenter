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
	"errors"
	"fmt"
)

const severityRuleMinConditionsCount = 1

type Severity struct {
	Default string         `mapstructure:"default"`
	Rules   []SeverityRule `mapstructure:"rules"`
}

func (s *Severity) Validate() error {
	if len(s.Rules) > 0 && s.Default == "" {
		return errors.New("can't set severity rule option: no default severity defined")
	}

	for i, rule := range s.Rules {
		if err := rule.Validate(); err != nil {
			return fmt.Errorf("error in severity rule #%d: %w", i, err)
		}
	}

	return nil
}

type SeverityRule struct {
	BaseRule `mapstructure:",squash"`
	Severity string `mapstructure:"severity"`
}

func (s *SeverityRule) Validate() error {
	if s.Severity == "" {
		return errors.New("severity should be set")
	}

	return s.BaseRule.Validate(severityRuleMinConditionsCount)
}

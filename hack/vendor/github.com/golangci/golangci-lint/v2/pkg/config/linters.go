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
	GroupStandard = "standard"
	GroupAll      = "all"
	GroupNone     = "none"
	GroupFast     = "fast"
)

type Linters struct {
	Default  string   `mapstructure:"default"`
	Enable   []string `mapstructure:"enable"`
	Disable  []string `mapstructure:"disable"`
	FastOnly bool     `mapstructure:"fast-only"` // Flag only option.

	Settings LintersSettings `mapstructure:"settings"`

	Exclusions LinterExclusions `mapstructure:"exclusions"`
}

func (l *Linters) Validate() error {
	validators := []func() error{
		l.Exclusions.Validate,
		l.validateNoFormatters,
	}

	for _, v := range validators {
		if err := v(); err != nil {
			return err
		}
	}

	return nil
}

func (l *Linters) validateNoFormatters() error {
	for _, n := range slices.Concat(l.Enable, l.Disable) {
		if slices.Contains(getAllFormatterNames(), n) {
			return fmt.Errorf("%s is a formatter", n)
		}
	}

	return nil
}

func getAllFormatterNames() []string {
	return []string{"gci", "gofmt", "gofumpt", "goimports", "golines", "swaggo"}
}

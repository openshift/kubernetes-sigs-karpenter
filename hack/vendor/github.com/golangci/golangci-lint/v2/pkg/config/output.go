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
	"strings"

	"github.com/golangci/golangci-lint/v2/pkg/fsutils"
)

type Output struct {
	Formats    Formats  `mapstructure:"formats"`
	SortOrder  []string `mapstructure:"sort-order"`
	ShowStats  bool     `mapstructure:"show-stats"`
	PathPrefix string   `mapstructure:"path-prefix"`
	PathMode   string   `mapstructure:"path-mode"`
}

func (o *Output) Validate() error {
	validators := []func() error{
		o.validateSortOrder,
		o.validatePathMode,
	}

	for _, v := range validators {
		if err := v(); err != nil {
			return err
		}
	}

	return nil
}

func (o *Output) validateSortOrder() error {
	validOrders := []string{"linter", "file", "severity"}

	all := strings.Join(o.SortOrder, " ")

	for _, order := range o.SortOrder {
		if strings.Count(all, order) > 1 {
			return fmt.Errorf("the sort-order name %q is repeated several times", order)
		}

		if !slices.Contains(validOrders, order) {
			return fmt.Errorf("unsupported sort-order name %q", order)
		}
	}

	return nil
}

func (o *Output) validatePathMode() error {
	switch o.PathMode {
	case "", fsutils.OutputPathModeAbsolute:
		// Valid

	default:
		return fmt.Errorf("unsupported output path mode %q", o.PathMode)
	}

	return nil
}

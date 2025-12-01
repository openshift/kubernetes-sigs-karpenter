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

type Formatters struct {
	Enable     []string            `mapstructure:"enable"`
	Settings   FormatterSettings   `mapstructure:"settings"`
	Exclusions FormatterExclusions `mapstructure:"exclusions"`
}

func (f *Formatters) Validate() error {
	for _, n := range f.Enable {
		if !slices.Contains(getAllFormatterNames(), n) {
			return fmt.Errorf("%s is not a formatter", n)
		}
	}

	return nil
}

type FormatterExclusions struct {
	Generated  string   `mapstructure:"generated"`
	Paths      []string `mapstructure:"paths"`
	WarnUnused bool     `mapstructure:"warn-unused"`
}

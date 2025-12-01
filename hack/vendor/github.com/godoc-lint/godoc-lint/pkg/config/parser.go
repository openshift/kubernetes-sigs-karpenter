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
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// FromYAML parses configuration from given YAML content.
func FromYAML(in []byte) (*PlainConfig, error) {
	raw := PlainConfig{}
	if err := yaml.Unmarshal(in, &raw); err != nil {
		return nil, fmt.Errorf("cannot parse config from YAML file: %w", err)
	}

	if raw.Version != nil && !strings.HasPrefix(*raw.Version, "1.") {
		return nil, fmt.Errorf("unsupported config version: %s", *raw.Version)
	}

	return &raw, nil
}

// FromYAMLFile parses configuration from given file path.
func FromYAMLFile(path string) (*PlainConfig, error) {
	in, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file (%s): %w", path, err)
	}

	raw := PlainConfig{}
	if err := yaml.Unmarshal(in, &raw); err != nil {
		return nil, fmt.Errorf("cannot parse config from YAML file: %w", err)
	}

	if raw.Version != nil && !strings.HasPrefix(*raw.Version, "1.") {
		return nil, fmt.Errorf("unsupported config version: %s", *raw.Version)
	}

	return &raw, nil
}

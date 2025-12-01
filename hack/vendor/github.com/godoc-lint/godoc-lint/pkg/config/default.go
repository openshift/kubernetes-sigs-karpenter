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
	_ "embed"
	"sync"
)

// defaultConfigFiles is the list of default configuration file names.
var defaultConfigFiles = []string{
	".godoc-lint.yaml",
	".godoc-lint.yml",
	".godoc-lint.json",
	".godoclint.yaml",
	".godoclint.yml",
	".godoclint.json",
}

// defaultConfigYAML is the default configuration (as YAML).
//
//go:embed default.yaml
var defaultConfigYAML []byte

// getDefaultPlainConfig returns the parsed default configuration.
var getDefaultPlainConfig = sync.OnceValue(func() *PlainConfig {
	// Error is nil due to tests.
	pcfg, _ := FromYAML(defaultConfigYAML)
	return pcfg
})

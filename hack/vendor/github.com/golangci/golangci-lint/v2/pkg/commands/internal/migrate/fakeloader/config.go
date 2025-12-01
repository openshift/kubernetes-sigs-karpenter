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

package fakeloader

// Config implements [config.BaseConfig].
// This only the stub for the real file loader.
type Config struct {
	Version string `mapstructure:"version"`

	cfgDir string // Path to the directory containing golangci-lint config file.
}

func NewConfig() *Config {
	return &Config{}
}

// SetConfigDir sets the path to directory that contains golangci-lint config file.
func (c *Config) SetConfigDir(dir string) {
	c.cfgDir = dir
}

func (*Config) IsInternalTest() bool {
	return false
}

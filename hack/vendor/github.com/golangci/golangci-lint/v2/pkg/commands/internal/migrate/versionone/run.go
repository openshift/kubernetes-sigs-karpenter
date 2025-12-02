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

package versionone

import (
	"time"
)

// Run encapsulates the config options for running the linter analysis.
type Run struct {
	Timeout time.Duration `mapstructure:"timeout"`

	Concurrency *int `mapstructure:"concurrency"`

	Go *string `mapstructure:"go"`

	RelativePathMode *string `mapstructure:"relative-path-mode"`

	BuildTags           []string `mapstructure:"build-tags"`
	ModulesDownloadMode *string  `mapstructure:"modules-download-mode"`

	ExitCodeIfIssuesFound *int  `mapstructure:"issues-exit-code"`
	AnalyzeTests          *bool `mapstructure:"tests"`

	AllowParallelRunners *bool `mapstructure:"allow-parallel-runners"`
	AllowSerialRunners   *bool `mapstructure:"allow-serial-runners"`
}

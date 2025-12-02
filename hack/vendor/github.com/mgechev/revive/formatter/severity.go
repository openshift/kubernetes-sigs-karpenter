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

package formatter

import "github.com/mgechev/revive/lint"

func severity(config lint.Config, failure lint.Failure) lint.Severity {
	if config, ok := config.Rules[failure.RuleName]; ok && config.Severity == lint.SeverityError {
		return lint.SeverityError
	}
	if config, ok := config.Directives[failure.RuleName]; ok && config.Severity == lint.SeverityError {
		return lint.SeverityError
	}
	return lint.SeverityWarning
}

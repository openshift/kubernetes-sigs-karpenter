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

import "github.com/butuzov/ireturn/analyzer/internal/types"

// rejectConfig specifies a list of interfaces (keywords, patterns and regular expressions)
// that are rejected by ireturn as valid to return, any non listed interface are allowed.
type rejectConfig struct {
	*defaultConfig
}

func rejectAll(patterns []string) *rejectConfig {
	return &rejectConfig{&defaultConfig{List: patterns}}
}

func (rc *rejectConfig) IsValid(i types.IFace) bool {
	return !rc.Has(i)
}

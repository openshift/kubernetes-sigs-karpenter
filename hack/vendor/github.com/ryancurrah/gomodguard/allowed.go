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

package gomodguard

import "strings"

// Allowed is a list of modules and module
// domains that are allowed to be used.
type Allowed struct {
	Modules []string `yaml:"modules"`
	Domains []string `yaml:"domains"`
}

// IsAllowedModule returns true if the given module
// name is in the allowed modules list.
func (a *Allowed) IsAllowedModule(moduleName string) bool {
	allowedModules := a.Modules

	for i := range allowedModules {
		if strings.TrimSpace(moduleName) == strings.TrimSpace(allowedModules[i]) {
			return true
		}
	}

	return false
}

// IsAllowedModuleDomain returns true if the given modules domain is
// in the allowed module domains list.
func (a *Allowed) IsAllowedModuleDomain(moduleName string) bool {
	allowedDomains := a.Domains

	for i := range allowedDomains {
		if strings.HasPrefix(strings.TrimSpace(strings.ToLower(moduleName)),
			strings.TrimSpace(strings.ToLower(allowedDomains[i]))) {
			return true
		}
	}

	return false
}

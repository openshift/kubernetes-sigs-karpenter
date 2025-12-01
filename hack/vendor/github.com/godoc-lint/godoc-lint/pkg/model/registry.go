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

package model

// Registry defines a registry of checkers.
type Registry interface {
	// Add registers a new checker.
	Add(Checker)

	// List returns a slice of the registered checkers.
	List() []Checker

	// GetCoveredRules returns the set of rules covered by the registered
	// checkers.
	GetCoveredRules() RuleSet
}

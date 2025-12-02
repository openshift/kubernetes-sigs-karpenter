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

// Package check provides a registry of checkers.
package check

import (
	"github.com/godoc-lint/godoc-lint/pkg/check/deprecated"
	"github.com/godoc-lint/godoc-lint/pkg/check/max_len"
	"github.com/godoc-lint/godoc-lint/pkg/check/no_unused_link"
	"github.com/godoc-lint/godoc-lint/pkg/check/pkg_doc"
	"github.com/godoc-lint/godoc-lint/pkg/check/require_doc"
	"github.com/godoc-lint/godoc-lint/pkg/check/start_with_name"
	"github.com/godoc-lint/godoc-lint/pkg/model"
)

// Registry implements a registry of rules.
type Registry struct {
	checkers     map[model.Checker]struct{}
	coveredRules model.RuleSet
}

// NewRegistry returns a new rule registry instance.
func NewRegistry(checkers ...model.Checker) *Registry {
	registry := Registry{
		checkers: make(map[model.Checker]struct{}, len(checkers)+10),
	}
	for _, c := range checkers {
		registry.Add(c)
	}
	return &registry
}

// NewPopulatedRegistry returns a registry with all supported rules registered.
func NewPopulatedRegistry() *Registry {
	return NewRegistry(
		max_len.NewMaxLenChecker(),
		pkg_doc.NewPkgDocChecker(),
		require_doc.NewRequireDocChecker(),
		start_with_name.NewStartWithNameChecker(),
		no_unused_link.NewNoUnusedLinkChecker(),
		deprecated.NewDeprecatedChecker(),
	)
}

// Add implements the corresponding interface method.
func (r *Registry) Add(checker model.Checker) {
	if _, ok := r.checkers[checker]; ok {
		return
	}
	r.coveredRules = r.coveredRules.Merge(checker.GetCoveredRules())
	r.checkers[checker] = struct{}{}
}

// List implements the corresponding interface method.
func (r *Registry) List() []model.Checker {
	all := make([]model.Checker, 0, len(r.checkers))
	for c := range r.checkers {
		all = append(all, c)
	}
	return all
}

// GetCoveredRules implements the corresponding interface method.
func (r *Registry) GetCoveredRules() model.RuleSet {
	return r.coveredRules
}

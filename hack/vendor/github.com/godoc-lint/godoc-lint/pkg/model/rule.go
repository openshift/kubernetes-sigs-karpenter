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

// Rule represents a rule.
type Rule string

const (
	// PkgDocRule represents the "pkg-doc" rule.
	PkgDocRule Rule = "pkg-doc"
	// SinglePkgDocRule represents the "single-pkg-doc" rule.
	SinglePkgDocRule Rule = "single-pkg-doc"
	// RequirePkgDocRule represents the "require-pkg-doc" rule.
	RequirePkgDocRule Rule = "require-pkg-doc"
	// StartWithNameRule represents the "start-with-name" rule.
	StartWithNameRule Rule = "start-with-name"
	// RequireDocRule represents the "require-doc" rule.
	RequireDocRule Rule = "require-doc"
	// DeprecatedRule represents the "deprecated" rule.
	DeprecatedRule Rule = "deprecated"
	// MaxLenRule represents the "max-len" rule.
	MaxLenRule Rule = "max-len"
	// NoUnusedLinkRule represents the "no-unused-link" rule.
	NoUnusedLinkRule Rule = "no-unused-link"
)

// AllRules is the set of all supported rules.
var AllRules = func() RuleSet {
	return RuleSet{}.Add(
		PkgDocRule,
		SinglePkgDocRule,
		RequirePkgDocRule,
		StartWithNameRule,
		RequireDocRule,
		DeprecatedRule,
		MaxLenRule,
		NoUnusedLinkRule,
	)
}()

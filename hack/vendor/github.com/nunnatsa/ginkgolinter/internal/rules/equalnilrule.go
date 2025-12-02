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

package rules

import (
	"github.com/nunnatsa/ginkgolinter/config"
	"github.com/nunnatsa/ginkgolinter/internal/expression"
	"github.com/nunnatsa/ginkgolinter/internal/expression/matcher"
	"github.com/nunnatsa/ginkgolinter/internal/reports"
)

// EqualNilRule checks for correct usage of nil comparisons.
// It suggests using the BeNil() matcher instead of Equal(nil) for better readability.
//
// Example:
//
//	// Bad:
//	Expect(x).To(Equal(nil))
//
//	// Good:
//	Expect(x).To(BeNil())
type EqualNilRule struct{}

func (r EqualNilRule) isApplied(gexp *expression.GomegaExpression, config config.Config) bool {
	return !config.SuppressNil &&
		gexp.MatcherTypeIs(matcher.EqualValueMatcherType)
}

func (r EqualNilRule) Apply(gexp *expression.GomegaExpression, config config.Config, reportBuilder *reports.Builder) bool {
	if !r.isApplied(gexp, config) {
		return false
	}

	gexp.SetMatcherBeNil()

	reportBuilder.AddIssue(true, wrongNilWarningTemplate)

	return true
}

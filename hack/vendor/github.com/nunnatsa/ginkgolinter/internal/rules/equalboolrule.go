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

const wrongBoolWarningTemplate = "wrong boolean assertion"

// EqualBoolRule checks for correct usage of boolean assertions.
// It suggests using the BeTrue() or BeFalse() matchers instead of Equal(true) or Equal(false).
//
// Example:
//
//	// Bad:
//	Expect(x).To(Equal(true))
//
//	// Good:
//	Expect(x).To(BeTrue())
type EqualBoolRule struct{}

func (r EqualBoolRule) isApplied(gexp *expression.GomegaExpression) bool {
	return gexp.MatcherTypeIs(matcher.EqualBoolValueMatcherType)
}

func (r EqualBoolRule) Apply(gexp *expression.GomegaExpression, _ config.Config, reportBuilder *reports.Builder) bool {
	if !r.isApplied(gexp) {
		return false
	}

	if gexp.MatcherTypeIs(matcher.BoolValueTrue) {
		gexp.SetMatcherBeTrue()
	} else {
		if gexp.IsNegativeAssertion() {
			gexp.ReverseAssertionFuncLogic()
			gexp.SetMatcherBeTrue()
		} else {
			gexp.SetMatcherBeFalse()
		}
	}

	reportBuilder.AddIssue(true, wrongBoolWarningTemplate)
	return true
}

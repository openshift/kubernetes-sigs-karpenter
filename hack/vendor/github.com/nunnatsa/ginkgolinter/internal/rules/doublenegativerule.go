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

const doubleNegativeWarningTemplate = "avoid double negative assertion"

// DoubleNegativeRule checks for double negative assertions, such as using `Not(BeFalse())`.
// It suggests replacing `Not(BeFalse())` with `BeTrue()` for better readability.
//
// Example:
//
//	// Bad:
//	Expect(x).NotTo(BeFalse())
//
//	// Good:
//	Expect(x).To(BeTrue())
type DoubleNegativeRule struct{}

func (DoubleNegativeRule) isApplied(gexp *expression.GomegaExpression) bool {
	return gexp.MatcherTypeIs(matcher.BeFalseMatcherType) &&
		gexp.IsNegativeAssertion()
}

func (r DoubleNegativeRule) Apply(gexp *expression.GomegaExpression, _ config.Config, reportBuilder *reports.Builder) bool {
	if !r.isApplied(gexp) {
		return false
	}

	gexp.ReverseAssertionFuncLogic()
	gexp.SetMatcherBeTrue()

	reportBuilder.AddIssue(true, doubleNegativeWarningTemplate)

	return true
}

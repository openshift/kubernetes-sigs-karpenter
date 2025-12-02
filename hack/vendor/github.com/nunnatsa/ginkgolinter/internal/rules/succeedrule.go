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
	"github.com/nunnatsa/ginkgolinter/internal/expression/actual"
	"github.com/nunnatsa/ginkgolinter/internal/expression/matcher"
	"github.com/nunnatsa/ginkgolinter/internal/reports"
)

// SucceedRule checks for correct usage of the Succeed matcher.
// It suggests using the HaveOccurred matcher for non-function error value, instead of Succeed.
//
// Example:
//
//	// Bad:
//	Expect(err).To(Succeed())
//
//	// Good:
//	Expect(err).ToNot(HaveOccurred())
//
// It also check that the actual value is a single value, and not a tuple.
//
// Example:
//
//	func doSomething() (string, error) {
//		return "hello", nil
//	}
//
//	// Bad:
//	Expect(doSomething()).To(Succeed())
//
//	// Good:
//	s, err := doSomething()
//	Expect(err).To(Succeed())
//
// In addition, it checks that the actual value is an error type.
//
// Example:
//
//	x := 5
//
//	// Bad:
//	Expect(x).To(Succeed())
type SucceedRule struct{}

func (r SucceedRule) isApplied(gexp *expression.GomegaExpression) bool {
	return !gexp.IsAsync() && gexp.MatcherTypeIs(matcher.SucceedMatcherType)
}

func (r SucceedRule) Apply(gexp *expression.GomegaExpression, config config.Config, reportBuilder *reports.Builder) bool {
	if !r.isApplied(gexp) {
		return false
	}

	if !gexp.ActualArgTypeIs(actual.ErrorTypeArgType) {
		if gexp.IsActualTuple() {
			reportBuilder.AddIssue(false, "the Success matcher does not support multiple values")
		} else {
			reportBuilder.AddIssue(false, "asserting a non-error type with Succeed matcher")
		}
		return true
	}

	if config.ForceSucceedForFuncs && !gexp.GetActualArg().(*actual.ErrPayload).IsFunc() {
		gexp.ReverseAssertionFuncLogic()
		gexp.SetMatcherHaveOccurred()

		reportBuilder.AddIssue(true, "prefer using the HaveOccurred matcher for non-function error value, instead of Succeed")

		return true
	}

	return false
}

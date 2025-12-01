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
	"go/token"

	"github.com/nunnatsa/ginkgolinter/config"
	"github.com/nunnatsa/ginkgolinter/internal/expression"
	"github.com/nunnatsa/ginkgolinter/internal/expression/actual"
	"github.com/nunnatsa/ginkgolinter/internal/expression/matcher"
	"github.com/nunnatsa/ginkgolinter/internal/reports"
)

const wrongCompareWarningTemplate = "wrong comparison assertion"

// ComparisonRule rewrites assertions that use comparison operators into more idiomatic matcher-based assertions.
//
// Examples:
//
//	Expect(x == 5).To(BeTrue())      // becomes: Expect(x).To(Equal(5))
//
//	Expect(y != 0).To(BeTrue())      // becomes: Expect(y).NotTo(BeZero())
//
//	Expect(a > b).To(BeTrue())       // becomes: Expect(a).To(BeNumerically(">", b))
//
//	Expect(count <= max).To(BeTrue()) // becomes: Expect(count).To(BeNumerically("<=", max))
type ComparisonRule struct{}

func (r ComparisonRule) isApplied(gexp *expression.GomegaExpression, config config.Config) bool {
	if config.SuppressCompare {
		return false
	}

	return gexp.ActualArgTypeIs(actual.ComparisonActualArgType)
}

func (r ComparisonRule) Apply(gexp *expression.GomegaExpression, config config.Config, reportBuilder *reports.Builder) bool {
	if !r.isApplied(gexp, config) {
		return false
	}

	actl, ok := gexp.GetActualArg().(actual.ComparisonActualPayload)
	if !ok {
		return false
	}

	switch actl.GetOp() {
	case token.EQL:
		r.handleEqualComparison(gexp, actl)

	case token.NEQ:
		gexp.ReverseAssertionFuncLogic()
		r.handleEqualComparison(gexp, actl)
	case token.GTR, token.GEQ, token.LSS, token.LEQ:
		if !actl.GetRight().IsValueNumeric() {
			return false
		}

		gexp.SetMatcherBeNumerically(actl.GetOp(), actl.GetRight().GetValueExpr())

	default:
		return false
	}

	if gexp.MatcherTypeIs(matcher.BoolValueFalse) {
		gexp.ReverseAssertionFuncLogic()
	}

	gexp.ReplaceActual(actl.GetLeft().GetValueExpr())

	reportBuilder.AddIssue(true, wrongCompareWarningTemplate)
	return true
}

func (r ComparisonRule) handleEqualComparison(gexp *expression.GomegaExpression, actual actual.ComparisonActualPayload) {
	if actual.GetRight().IsValueZero() {
		gexp.SetMatcherBeZero()
	} else {
		left := actual.GetLeft()
		arg := actual.GetRight().GetValueExpr()
		if left.IsInterface() || left.IsPointer() {
			gexp.SetMatcherBeIdenticalTo(arg)
		} else {
			gexp.SetMatcherEqual(arg)
		}
	}
}

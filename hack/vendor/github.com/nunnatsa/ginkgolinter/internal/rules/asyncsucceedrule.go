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

// AsyncSucceedRule ensures that the Succeed matcher is used correctly with asynchronous functions.
// It flags cases where the function returns multiple values, or when the function does not return a single error value
// or does not have Gomega as its first parameter, as these usages are not supported by the Succeed matcher.
type AsyncSucceedRule struct{}

func (AsyncSucceedRule) isApply(gexp *expression.GomegaExpression) bool {
	return gexp.IsAsync() &&
		gexp.MatcherTypeIs(matcher.SucceedMatcherType) &&
		gexp.ActualArgTypeIs(actual.FuncSigArgType) &&
		!gexp.ActualArgTypeIs(actual.ErrorTypeArgType|actual.GomegaParamArgType)
}

func (r AsyncSucceedRule) Apply(gexp *expression.GomegaExpression, _ config.Config, reportBuilder *reports.Builder) bool {
	if r.isApply(gexp) {
		if gexp.ActualArgTypeIs(actual.MultiRetsArgType) {
			reportBuilder.AddIssue(false, "Success matcher does not support multiple values")
		} else {
			reportBuilder.AddIssue(false, "Success matcher only support a single error value, or function with Gomega as its first parameter")
		}
	}

	return false
}

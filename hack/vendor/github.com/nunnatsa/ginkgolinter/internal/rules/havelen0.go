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

// HaveLen0 rule checks if the HaveLen matcher is used with zero.
// It suggests using the BeEmpty() matcher instead of HaveLen(0) for better readability.
//
// Example:
//
//	var s []string
//
//	// Bad:
//	Expect(s).To(HaveLen(0))
//
//	// Good:
//	Expect(s).To(BeEmpty())
type HaveLen0 struct{}

func (r *HaveLen0) isApplied(gexp *expression.GomegaExpression, config config.Config) bool {
	return gexp.MatcherTypeIs(matcher.HaveLenZeroMatcherType) && !config.AllowHaveLen0
}

func (r *HaveLen0) Apply(gexp *expression.GomegaExpression, config config.Config, reportBuilder *reports.Builder) bool {
	if !r.isApplied(gexp, config) {
		return false
	}
	gexp.SetMatcherBeEmpty()
	reportBuilder.AddIssue(true, wrongLengthWarningTemplate)
	return true
}

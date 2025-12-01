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
	"github.com/nunnatsa/ginkgolinter/internal/reports"
)

const forceExpectToTemplate = "must not use %s with %s"

// ForceExpectToRule checks for correct usage of Expect and ExpectWithOffset assertions.
// It suggests using the To and ToNot assertion methods instead of the Should and ShouldNot.
//
// Example:
//
//	// Bad:
//	Expect(x).Should(Equal(5))
//
//	// Good:
//	Expect(x).To(Equal(5))
type ForceExpectToRule struct{}

func (ForceExpectToRule) isApplied(gexp *expression.GomegaExpression, config config.Config) bool {
	if !config.ForceExpectTo {
		return false
	}

	actlName := gexp.GetActualFuncName()
	return actlName == "Expect" || actlName == "ExpectWithOffset"
}

func (r ForceExpectToRule) Apply(gexp *expression.GomegaExpression, config config.Config, reportBuilder *reports.Builder) bool {
	if !r.isApplied(gexp, config) {
		return false
	}

	var newName string

	switch gexp.GetAssertFuncName() {
	case "Should":
		newName = "To"
	case "ShouldNot":
		newName = "ToNot"
	default:
		return false
	}

	gexp.ReplaceAssertionMethod(newName)
	reportBuilder.AddIssue(true, forceExpectToTemplate, gexp.GetActualFuncName(), gexp.GetOrigAssertFuncName())

	// always return false, to keep checking another rules.
	return false
}

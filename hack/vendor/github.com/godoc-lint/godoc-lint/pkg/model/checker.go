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

import "golang.org/x/tools/go/analysis"

// AnalysisContext provides contextual information about the running analysis.
type AnalysisContext struct {
	// Config provides analyzer configuration.
	Config Config

	// InspectorResult is the analysis result of the pre-run inspector.
	InspectorResult *InspectorResult

	// Pass is the analysis Pass instance.
	Pass *analysis.Pass
}

// Checker defines a rule checker.
type Checker interface {
	// GetCoveredRules returns the set of rules applied by the checker.
	GetCoveredRules() RuleSet

	// Apply checks for the rule(s).
	Apply(actx *AnalysisContext) error
}

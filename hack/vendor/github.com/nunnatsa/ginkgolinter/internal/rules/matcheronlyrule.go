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

// matcherOnlyRules is a collection of rules that only validate the matcher part of the assertion.
// It does not validate the actual part of the assertion.
var matcherOnlyRules = Rules{
	&HaveLen0{},
	&EqualBoolRule{},
	&EqualNilRule{},
	&DoubleNegativeRule{},
	// must be the last rule in the list
	&SimplifyNotRule{},
}

func getMatcherOnlyRules() Rules {
	return matcherOnlyRules
}

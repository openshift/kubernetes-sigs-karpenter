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

package flect

var singularRules = []rule{}

// AddSingular adds a rule that will replace the given suffix with the replacement suffix.
// The name is confusing. This function will be deprecated in the next release.
func AddSingular(ext string, repl string) {
	InsertSingularRule(ext, repl)
}

// InsertSingularRule inserts a rule that will replace the given suffix with
// the repl(acement) at the beginning of the list of the singularize rules.
func InsertSingularRule(suffix, repl string) {
	singularMoot.Lock()
	defer singularMoot.Unlock()

	singularRules = append([]rule{{
		suffix: suffix,
		fn:     simpleRuleFunc(suffix, repl),
	}}, singularRules...)

	singularRules = append([]rule{{
		suffix: repl,
		fn:     noop,
	}}, singularRules...)
}

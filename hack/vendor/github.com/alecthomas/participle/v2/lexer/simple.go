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

package lexer

// SimpleRule is a named regular expression.
type SimpleRule struct {
	Name    string
	Pattern string
}

// MustSimple creates a new Stateful lexer with only a single root state.
// The rules are tried in order.
//
// It panics if there is an error.
func MustSimple(rules []SimpleRule) *StatefulDefinition {
	def, err := NewSimple(rules)
	if err != nil {
		panic(err)
	}
	return def
}

// NewSimple creates a new Stateful lexer with only a single root state.
// The rules are tried in order.
func NewSimple(rules []SimpleRule) (*StatefulDefinition, error) {
	fullRules := make([]Rule, len(rules))
	for i, rule := range rules {
		fullRules[i] = Rule{Name: rule.Name, Pattern: rule.Pattern}
	}
	return New(Rules{"Root": fullRules})
}

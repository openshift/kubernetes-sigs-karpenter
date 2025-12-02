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

package yqlib

import (
	"container/list"
	"fmt"
	"strconv"
)

func tryConvertToNumber(value string) (string, bool) {
	// try an int first
	_, _, err := parseInt64(value)
	if err == nil {
		return "!!int", true
	}
	// try float
	_, floatErr := strconv.ParseFloat(value, 64)

	if floatErr == nil {
		return "!!float", true
	}
	return "", false

}

func toNumberOperator(_ *dataTreeNavigator, context Context, _ *ExpressionNode) (Context, error) {
	log.Debugf("ToNumberOperator")

	var results = list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)
		if candidate.Kind != ScalarNode {
			return Context{}, fmt.Errorf("cannot convert node at path %v of tag %v to number", candidate.GetNicePath(), candidate.Tag)
		}

		if candidate.Tag == "!!int" || candidate.Tag == "!!float" {
			// it already is a number!
			results.PushBack(candidate)
		} else {
			tag, converted := tryConvertToNumber(candidate.Value)
			if converted {
				result := candidate.CreateReplacement(ScalarNode, tag, candidate.Value)
				results.PushBack(result)
			} else {
				return Context{}, fmt.Errorf("cannot convert node value [%v] at path %v of tag %v to number", candidate.Value, candidate.GetNicePath(), candidate.Tag)
			}

		}
	}

	return context.ChildContext(results), nil
}

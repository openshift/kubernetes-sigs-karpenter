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

import "container/list"

type parentOpPreferences struct {
	Level int
}

func getParentOperator(_ *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {
	log.Debugf("getParentOperator")

	var results = list.New()

	prefs := expressionNode.Operation.Preferences.(parentOpPreferences)

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)
		currentLevel := 0
		for currentLevel < prefs.Level && candidate != nil {
			log.Debugf("currentLevel: %v, desired: %v", currentLevel, prefs.Level)
			log.Debugf("candidate: %v", NodeToString(candidate))
			candidate = candidate.Parent
			currentLevel++
		}

		log.Debugf("found candidate: %v", NodeToString(candidate))
		if candidate != nil {
			results.PushBack(candidate)
		}
	}

	return context.ChildContext(results), nil

}

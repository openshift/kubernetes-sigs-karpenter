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
)

type recursiveDescentPreferences struct {
	TraversePreferences traversePreferences
	RecurseArray        bool
}

func recursiveDescentOperator(_ *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {
	var results = list.New()

	preferences := expressionNode.Operation.Preferences.(recursiveDescentPreferences)
	err := recursiveDecent(results, context, preferences)
	if err != nil {
		return Context{}, err
	}

	return context.ChildContext(results), nil
}

func recursiveDecent(results *list.List, context Context, preferences recursiveDescentPreferences) error {
	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)

		log.Debugf("added %v", NodeToString(candidate))
		results.PushBack(candidate)

		if candidate.Kind != AliasNode && len(candidate.Content) > 0 &&
			(preferences.RecurseArray || candidate.Kind != SequenceNode) {

			children, err := splat(context.SingleChildContext(candidate), preferences.TraversePreferences)

			if err != nil {
				return err
			}
			err = recursiveDecent(results, children, preferences)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

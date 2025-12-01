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

func selectOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {

	log.Debugf("selectOperation")
	var results = list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)
		rhs, err := d.GetMatchingNodes(context.SingleReadonlyChildContext(candidate), expressionNode.RHS)

		if err != nil {
			return Context{}, err
		}

		// find any truthy node
		includeResult := false

		for resultEl := rhs.MatchingNodes.Front(); resultEl != nil; resultEl = resultEl.Next() {
			result := resultEl.Value.(*CandidateNode)
			includeResult = isTruthyNode(result)
			if includeResult {
				break
			}
		}

		if includeResult {
			results.PushBack(candidate)
		}
	}
	return context.ChildContext(results), nil
}

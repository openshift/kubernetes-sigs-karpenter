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

func referenceOperator(_ *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {
	return context.SingleChildContext(expressionNode.Operation.CandidateNode), nil
}

func valueOperator(_ *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {
	log.Debug("value = %v", expressionNode.Operation.CandidateNode.Value)
	if context.MatchingNodes.Len() == 0 {
		clone := expressionNode.Operation.CandidateNode.Copy()
		return context.SingleChildContext(clone), nil
	}

	var results = list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		clone := expressionNode.Operation.CandidateNode.Copy()
		results.PushBack(clone)
	}

	return context.ChildContext(results), nil
}

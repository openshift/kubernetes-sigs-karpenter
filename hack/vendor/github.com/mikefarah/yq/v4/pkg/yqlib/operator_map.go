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

func mapValuesOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)
		//run expression against entries
		// splat toEntries and pipe it into Rhs
		splatted, err := splat(context.SingleChildContext(candidate), traversePreferences{})
		if err != nil {
			return Context{}, err
		}

		assignUpdateExp := &ExpressionNode{
			Operation: &Operation{OperationType: assignOpType, UpdateAssign: true},
			RHS:       expressionNode.RHS,
		}
		_, err = assignUpdateOperator(d, splatted, assignUpdateExp)
		if err != nil {
			return Context{}, err
		}

	}

	return context, nil
}

func mapOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {

	var results = list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)
		//run expression against entries
		// splat toEntries and pipe it into Rhs
		splatted, err := splat(context.SingleChildContext(candidate), traversePreferences{})
		if err != nil {
			return Context{}, err
		}
		if splatted.MatchingNodes.Len() == 0 {
			results.PushBack(candidate.Copy())
			continue
		}

		result, err := d.GetMatchingNodes(splatted, expressionNode.RHS)
		log.Debug("expressionNode.Rhs %v", expressionNode.RHS.Operation.OperationType)
		log.Debug("result %v", result)
		if err != nil {
			return Context{}, err
		}

		selfExpression := &ExpressionNode{Operation: &Operation{OperationType: selfReferenceOpType}}
		collected, err := collectTogether(d, result, selfExpression)
		if err != nil {
			return Context{}, err
		}
		collected.Style = candidate.Style

		results.PushBack(collected)

	}

	return context.ChildContext(results), nil
}

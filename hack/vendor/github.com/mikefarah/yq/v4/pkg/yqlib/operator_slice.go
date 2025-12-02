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
)

func getSliceNumber(d *dataTreeNavigator, context Context, node *CandidateNode, expressionNode *ExpressionNode) (int, error) {
	result, err := d.GetMatchingNodes(context.SingleChildContext(node), expressionNode)
	if err != nil {
		return 0, err
	}
	if result.MatchingNodes.Len() != 1 {
		return 0, fmt.Errorf("expected to find 1 number, got %v instead", result.MatchingNodes.Len())
	}
	return parseInt(result.MatchingNodes.Front().Value.(*CandidateNode).Value)
}

func sliceArrayOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {

	log.Debug("slice array operator!")
	log.Debug("lhs: %v", expressionNode.LHS.Operation.toString())
	log.Debug("rhs: %v", expressionNode.RHS.Operation.toString())

	results := list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		lhsNode := el.Value.(*CandidateNode)

		firstNumber, err := getSliceNumber(d, context, lhsNode, expressionNode.LHS)

		if err != nil {
			return Context{}, err
		}
		relativeFirstNumber := firstNumber
		if relativeFirstNumber < 0 {
			relativeFirstNumber = len(lhsNode.Content) + firstNumber
		}

		secondNumber, err := getSliceNumber(d, context, lhsNode, expressionNode.RHS)
		if err != nil {
			return Context{}, err
		}

		relativeSecondNumber := secondNumber
		if relativeSecondNumber < 0 {
			relativeSecondNumber = len(lhsNode.Content) + secondNumber
		} else if relativeSecondNumber > len(lhsNode.Content) {
			relativeSecondNumber = len(lhsNode.Content)
		}

		log.Debug("calculateIndicesToTraverse: slice from %v to %v", relativeFirstNumber, relativeSecondNumber)

		var newResults []*CandidateNode
		for i := relativeFirstNumber; i < relativeSecondNumber; i++ {
			newResults = append(newResults, lhsNode.Content[i])
		}

		sliceArrayNode := lhsNode.CreateReplacement(SequenceNode, lhsNode.Tag, "")
		sliceArrayNode.AddChildren(newResults)
		results.PushBack(sliceArrayNode)

	}

	// result is now the context that has the nodes we need to put back into a sequence.
	//what about multiple arrays in the context? I think we need to create an array for each one
	return context.ChildContext(results), nil
}

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

	"github.com/elliotchance/orderedmap"
)

func unique(d *dataTreeNavigator, context Context, _ *ExpressionNode) (Context, error) {
	selfExpression := &ExpressionNode{Operation: &Operation{OperationType: selfReferenceOpType}}
	uniqueByExpression := &ExpressionNode{Operation: &Operation{OperationType: uniqueByOpType}, RHS: selfExpression}
	return uniqueBy(d, context, uniqueByExpression)

}

func uniqueBy(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {

	log.Debugf("uniqueBy Operator")
	var results = list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)

		if candidate.Kind != SequenceNode {
			return Context{}, fmt.Errorf("only arrays are supported for unique")
		}

		var newMatches = orderedmap.NewOrderedMap()
		for _, child := range candidate.Content {
			rhs, err := d.GetMatchingNodes(context.SingleReadonlyChildContext(child), expressionNode.RHS)

			if err != nil {
				return Context{}, err
			}

			keyValue, err := getUniqueKeyValue(rhs)
			if err != nil {
				return Context{}, err
			}

			_, exists := newMatches.Get(keyValue)

			if !exists {
				newMatches.Set(keyValue, child)
			}
		}
		resultNode := candidate.CreateReplacementWithComments(SequenceNode, "!!seq", candidate.Style)
		for el := newMatches.Front(); el != nil; el = el.Next() {
			resultNode.AddChild(el.Value.(*CandidateNode))
		}

		results.PushBack(resultNode)
	}

	return context.ChildContext(results), nil

}

func getUniqueKeyValue(rhs Context) (string, error) {
	keyValue := "null"
	var err error

	if rhs.MatchingNodes.Len() > 0 {
		first := rhs.MatchingNodes.Front()
		keyCandidate := first.Value.(*CandidateNode)
		keyValue = keyCandidate.Value
		if keyCandidate.Kind != ScalarNode {
			keyValue, err = encodeToString(keyCandidate, encoderPreferences{YamlFormat, 0})
		}
	}
	return keyValue, err
}

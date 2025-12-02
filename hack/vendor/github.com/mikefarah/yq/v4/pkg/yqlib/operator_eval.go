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

func evalOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {
	log.Debugf("Eval")
	pathExpStrResults, err := d.GetMatchingNodes(context.ReadOnlyClone(), expressionNode.RHS)
	if err != nil {
		return Context{}, err
	}

	expressions := make([]*ExpressionNode, pathExpStrResults.MatchingNodes.Len())
	expIndex := 0
	//parse every expression
	for pathExpStrEntry := pathExpStrResults.MatchingNodes.Front(); pathExpStrEntry != nil; pathExpStrEntry = pathExpStrEntry.Next() {
		expressionStrCandidate := pathExpStrEntry.Value.(*CandidateNode)

		expressions[expIndex], err = ExpressionParser.ParseExpression(expressionStrCandidate.Value)
		if err != nil {
			return Context{}, err
		}

		expIndex++
	}

	results := list.New()

	for matchEl := context.MatchingNodes.Front(); matchEl != nil; matchEl = matchEl.Next() {
		for expIndex = 0; expIndex < len(expressions); expIndex++ {
			result, err := d.GetMatchingNodes(context, expressions[expIndex])
			if err != nil {
				return Context{}, err
			}
			results.PushBackList(result.MatchingNodes)
		}
	}

	return context.ChildContext(results), nil

}

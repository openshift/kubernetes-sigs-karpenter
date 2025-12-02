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

func reverseOperator(_ *dataTreeNavigator, context Context, _ *ExpressionNode) (Context, error) {
	results := list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)

		if candidate.Kind != SequenceNode {
			return context, fmt.Errorf("node at path [%v] is not an array (it's a %v)", candidate.GetNicePath(), candidate.Tag)
		}

		reverseList := candidate.CreateReplacementWithComments(SequenceNode, "!!seq", candidate.Style)
		reverseContent := make([]*CandidateNode, len(candidate.Content))

		for i, originalNode := range candidate.Content {
			reverseContent[len(candidate.Content)-i-1] = originalNode
		}
		reverseList.AddChildren(reverseContent)
		results.PushBack(reverseList)

	}

	return context.ChildContext(results), nil

}

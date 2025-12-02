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

func isKeyOperator(_ *dataTreeNavigator, context Context, _ *ExpressionNode) (Context, error) {
	log.Debugf("isKeyOperator")

	var results = list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)

		results.PushBack(createBooleanCandidate(candidate, candidate.IsMapKey))
	}

	return context.ChildContext(results), nil
}

func getKeyOperator(_ *dataTreeNavigator, context Context, _ *ExpressionNode) (Context, error) {
	log.Debugf("getKeyOperator")

	var results = list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)

		if candidate.Key != nil {
			results.PushBack(candidate.Key)
		}
	}

	return context.ChildContext(results), nil

}

func keysOperator(_ *dataTreeNavigator, context Context, _ *ExpressionNode) (Context, error) {
	log.Debugf("keysOperator")

	var results = list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)

		var targetNode *CandidateNode
		switch candidate.Kind {
		case MappingNode:
			targetNode = getMapKeys(candidate)
		case SequenceNode:
			targetNode = getIndices(candidate)
		default:
			return Context{}, fmt.Errorf("cannot get keys of %v, keys only works for maps and arrays", candidate.Tag)
		}

		results.PushBack(targetNode)
	}

	return context.ChildContext(results), nil
}

func getMapKeys(node *CandidateNode) *CandidateNode {
	contents := make([]*CandidateNode, 0)
	for index := 0; index < len(node.Content); index = index + 2 {
		contents = append(contents, node.Content[index])
	}
	return &CandidateNode{Kind: SequenceNode, Tag: "!!seq", Content: contents}
}

func getIndices(node *CandidateNode) *CandidateNode {
	var contents = make([]*CandidateNode, len(node.Content))

	for index := range node.Content {
		contents[index] = &CandidateNode{
			Kind:  ScalarNode,
			Tag:   "!!int",
			Value: fmt.Sprintf("%v", index),
		}
	}

	return &CandidateNode{Kind: SequenceNode, Tag: "!!seq", Content: contents}
}

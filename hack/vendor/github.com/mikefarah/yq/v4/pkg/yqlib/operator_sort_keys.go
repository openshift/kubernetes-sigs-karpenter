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
	"sort"
)

func sortKeysOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)
		rhs, err := d.GetMatchingNodes(context.SingleReadonlyChildContext(candidate), expressionNode.RHS)
		if err != nil {
			return Context{}, err
		}

		for childEl := rhs.MatchingNodes.Front(); childEl != nil; childEl = childEl.Next() {
			node := childEl.Value.(*CandidateNode)
			if node.Kind == MappingNode {
				sortKeys(node)
			}
			if err != nil {
				return Context{}, err
			}
		}

	}
	return context, nil
}

func sortKeys(node *CandidateNode) {
	keys := make([]string, len(node.Content)/2)
	keyBucket := map[string]*CandidateNode{}
	valueBucket := map[string]*CandidateNode{}
	var contents = node.Content
	for index := 0; index < len(contents); index = index + 2 {
		key := contents[index]
		value := contents[index+1]
		keys[index/2] = key.Value
		keyBucket[key.Value] = key
		valueBucket[key.Value] = value
	}
	sort.Strings(keys)
	sortedContent := make([]*CandidateNode, len(node.Content))
	for index := 0; index < len(keys); index = index + 1 {
		keyString := keys[index]
		sortedContent[index*2] = keyBucket[keyString]
		sortedContent[1+(index*2)] = valueBucket[keyString]
	}

	// re-arranging children, no need to update their parent
	// relationship
	node.Content = sortedContent
}

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
	"math/rand"
)

func shuffleOperator(_ *dataTreeNavigator, context Context, _ *ExpressionNode) (Context, error) {

	// ignore CWE-338 gosec issue of not using crypto/rand
	// this is just to shuffle an array rather generating a
	// secret or something that needs proper rand.
	myRand := rand.New(rand.NewSource(Now().UnixNano())) // #nosec

	results := list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)

		if candidate.Kind != SequenceNode {
			return context, fmt.Errorf("node at path [%v] is not an array (it's a %v)", candidate.GetNicePath(), candidate.Tag)
		}

		result := candidate.Copy()

		a := result.Content

		myRand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

		results.PushBack(result)
	}
	return context.ChildContext(results), nil
}

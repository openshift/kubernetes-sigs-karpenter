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

import "fmt"

func withOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {
	log.Debugf("withOperator")
	// with(path, exp)

	if expressionNode.RHS.Operation.OperationType != blockOpType {
		return Context{}, fmt.Errorf("with must be given a block (;), got %v instead", expressionNode.RHS.Operation.OperationType.Type)
	}

	pathExp := expressionNode.RHS.LHS

	updateContext, err := d.GetMatchingNodes(context, pathExp)

	if err != nil {
		return Context{}, err
	}

	updateExp := expressionNode.RHS.RHS

	for el := updateContext.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)
		_, err = d.GetMatchingNodes(updateContext.SingleChildContext(candidate), updateExp)
		if err != nil {
			return Context{}, err
		}

	}

	return context, nil

}

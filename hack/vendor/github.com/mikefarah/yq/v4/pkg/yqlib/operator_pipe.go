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

func pipeOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {

	if expressionNode.LHS.Operation.OperationType == assignVariableOpType {
		return variableLoop(d, context, expressionNode)
	}
	lhs, err := d.GetMatchingNodes(context, expressionNode.LHS)
	if err != nil {
		return Context{}, err
	}
	rhsContext := context.ChildContext(lhs.MatchingNodes)
	rhs, err := d.GetMatchingNodes(rhsContext, expressionNode.RHS)
	if err != nil {
		return Context{}, err
	}
	return context.ChildContext(rhs.MatchingNodes), nil
}

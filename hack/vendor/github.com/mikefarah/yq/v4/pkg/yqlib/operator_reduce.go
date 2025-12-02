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

func reduceOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {
	log.Debugf("reduceOp")
	//.a as $var reduce (0; . + $var)
	//lhs is the assignment operator
	//rhs is the reduce block
	// '.' refers to the current accumulator, initialised to 0
	// $var references a single element from the .a

	//ensure lhs is actually an assignment
	//and rhs is a block (empty)
	if expressionNode.LHS.Operation.OperationType != assignVariableOpType {
		return Context{}, fmt.Errorf("reduce must be given a variables assignment, got %v instead", expressionNode.LHS.Operation.OperationType.Type)
	} else if expressionNode.RHS.Operation.OperationType != blockOpType {
		return Context{}, fmt.Errorf("reduce must be given a block, got %v instead", expressionNode.RHS.Operation.OperationType.Type)
	}

	arrayExpNode := expressionNode.LHS.LHS
	array, err := d.GetMatchingNodes(context, arrayExpNode)

	if err != nil {
		return Context{}, err
	}

	variableName := expressionNode.LHS.RHS.Operation.StringValue

	initExp := expressionNode.RHS.LHS

	accum, err := d.GetMatchingNodes(context, initExp)
	if err != nil {
		return Context{}, err
	}

	log.Debugf("with variable %v", variableName)

	blockExp := expressionNode.RHS.RHS
	for el := array.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)
		log.Debugf("REDUCING WITH %v", NodeToString(candidate))
		l := list.New()
		l.PushBack(candidate)
		accum.SetVariable(variableName, l)

		accum, err = d.GetMatchingNodes(accum, blockExp)
		if err != nil {
			return Context{}, err
		}
	}

	return accum, nil
}

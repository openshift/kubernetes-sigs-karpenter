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

func alternativeOperator(d *dataTreeNavigator, context Context, expressionNode *ExpressionNode) (Context, error) {
	log.Debugf("alternative")
	prefs := crossFunctionPreferences{
		CalcWhenEmpty: true,
		Calculation:   alternativeFunc,
		LhsResultValue: func(lhs *CandidateNode) (*CandidateNode, error) {
			if lhs == nil {
				return nil, nil
			}
			truthy := isTruthyNode(lhs)
			if truthy {
				return lhs, nil
			}
			return nil, nil
		},
	}
	return crossFunctionWithPrefs(d, context, expressionNode, prefs)
}

func alternativeFunc(_ *dataTreeNavigator, _ Context, lhs *CandidateNode, rhs *CandidateNode) (*CandidateNode, error) {
	if lhs == nil {
		return rhs, nil
	}
	if rhs == nil {
		return lhs, nil
	}

	isTrue := isTruthyNode(lhs)
	if isTrue {
		return lhs, nil
	}
	return rhs, nil
}

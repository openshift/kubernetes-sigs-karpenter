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

package reverseassertion

import "go/token"

var reverseLogicAssertions = map[string]string{
	"To":        "ToNot",
	"ToNot":     "To",
	"NotTo":     "To",
	"Should":    "ShouldNot",
	"ShouldNot": "Should",
}

// ChangeAssertionLogic get gomega assertion function name, and returns the reverse logic function name
func ChangeAssertionLogic(funcName string) string {
	if revFunc, ok := reverseLogicAssertions[funcName]; ok {
		return revFunc
	}
	return funcName
}

func IsNegativeLogic(funcName string) bool {
	switch funcName {
	case "ToNot", "NotTo", "ShouldNot":
		return true
	}
	return false
}

var reverseCompareOperators = map[token.Token]token.Token{
	token.LSS: token.GTR,
	token.GTR: token.LSS,
	token.LEQ: token.GEQ,
	token.GEQ: token.LEQ,
}

// ChangeCompareOperator return the reversed comparison operator
func ChangeCompareOperator(op token.Token) token.Token {
	if revOp, ok := reverseCompareOperators[op]; ok {
		return revOp
	}
	return op
}

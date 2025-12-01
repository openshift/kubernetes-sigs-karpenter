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

package checkers

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func xor(a, b bool) bool {
	return a != b
}

// anyVal returns the first value[i] for which bools[i] is true.
func anyVal[T any](bools []bool, vals ...T) (T, bool) {
	if len(bools) != len(vals) {
		panic("inconsistent usage of valOr") //nolint:forbidigo // Does not depend on the code being analyzed.
	}

	for i, b := range bools {
		if b {
			return vals[i], true
		}
	}

	var _default T
	return _default, false
}

func anyCondSatisfaction(pass *analysis.Pass, p predicate, vals ...ast.Expr) bool {
	for _, v := range vals {
		if p(pass, v) {
			return true
		}
	}
	return false
}

// p transforms simple is-function in a predicate.
func p(fn func(e ast.Expr) bool) predicate {
	return func(_ *analysis.Pass, e ast.Expr) bool {
		return fn(e)
	}
}

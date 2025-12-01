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

package irutil

import (
	"honnef.co/go/tools/go/ir"
)

// IsStub reports whether a function is a stub. A function is
// considered a stub if it has no instructions or if all it does is
// return a constant value.
func IsStub(fn *ir.Function) bool {
	for _, b := range fn.Blocks {
		for _, instr := range b.Instrs {
			switch instr.(type) {
			case *ir.Const:
				// const naturally has no side-effects
			case *ir.Panic:
				// panic is a stub if it only uses constants
			case *ir.Return:
				// return is a stub if it only uses constants
			case *ir.DebugRef:
			case *ir.Jump:
				// if there are no disallowed instructions, then we're
				// only jumping to the exit block (or possibly
				// somewhere else that's stubby?)
			default:
				// all other instructions are assumed to do actual work
				return false
			}
		}
	}
	return true
}

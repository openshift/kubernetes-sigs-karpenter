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

package ssainstr

import (
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/ssa"
)

// LookupCalled looks up ssa.Instruction that call the `fn` func in the given instr
func LookupCalled(instr ssa.Instruction, fn *types.Func) ([]ssa.Instruction, bool) {
	instrs := []ssa.Instruction{}

	call, ok := instr.(ssa.CallInstruction)
	if !ok {
		return instrs, false
	}

	ssaCall := call.Value()
	if ssaCall == nil {
		return instrs, false
	}
	common := ssaCall.Common()
	if common == nil {
		return instrs, false
	}
	val := common.Value

	called := false
	switch fnval := val.(type) {
	case *ssa.Function:
		for _, block := range fnval.Blocks {
			for _, instr := range block.Instrs {
				if analysisutil.Called(instr, nil, fn) {
					called = true
					instrs = append(instrs, instr)
				}
			}
		}
	}

	return instrs, called
}

// HasArgs returns whether the given ssa.Instruction has `typ` type args
func HasArgs(instr ssa.Instruction, typ types.Type) bool {
	call, ok := instr.(ssa.CallInstruction)
	if !ok {
		return false
	}

	ssaCall := call.Value()
	if ssaCall == nil {
		return false
	}

	for _, arg := range ssaCall.Call.Args {
		if types.Identical(arg.Type(), typ) {
			return true
		}
	}
	return false
}

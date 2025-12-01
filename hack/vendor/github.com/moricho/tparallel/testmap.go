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

package tparallel

import (
	"go/types"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"

	"github.com/moricho/tparallel/pkg/ssainstr"
)

// getTestMap gets a set of a top-level test and its sub-tests
func getTestMap(ssaanalyzer *buildssa.SSA, testTyp types.Type) map[*ssa.Function][]*ssa.Function {
	testMap := map[*ssa.Function][]*ssa.Function{}

	trun := analysisutil.MethodOf(testTyp, "Run")
	for _, f := range ssaanalyzer.SrcFuncs {
		if !strings.HasPrefix(f.Name(), "Test") || !(f.Parent() == (*ssa.Function)(nil)) {
			continue
		}
		testMap[f] = []*ssa.Function{}
		for _, block := range f.Blocks {
			for _, instr := range block.Instrs {
				called := analysisutil.Called(instr, nil, trun)

				if !called && ssainstr.HasArgs(instr, types.NewPointer(testTyp)) {
					if instrs, ok := ssainstr.LookupCalled(instr, trun); ok {
						for _, v := range instrs {
							testMap[f] = appendTestMap(testMap[f], v)
						}
					}
				} else if called {
					testMap[f] = appendTestMap(testMap[f], instr)
				}
			}
		}
	}

	return testMap
}

// appendTestMap converts ssa.Instruction to ssa.Function and append it to a given sub-test slice
func appendTestMap(subtests []*ssa.Function, instr ssa.Instruction) []*ssa.Function {
	call, ok := instr.(ssa.CallInstruction)
	if !ok {
		return subtests
	}

	ssaCall := call.Value()
	if ssaCall == nil {
		return subtests
	}

	for _, arg := range ssaCall.Call.Args {
		switch arg := arg.(type) {
		case *ssa.Function:
			subtests = append(subtests, arg)
		case *ssa.MakeClosure:
			fn, _ := arg.Fn.(*ssa.Function)
			subtests = append(subtests, fn)
		}
	}

	return subtests
}

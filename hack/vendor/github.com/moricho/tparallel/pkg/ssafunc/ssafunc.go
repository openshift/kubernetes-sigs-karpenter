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

package ssafunc

import (
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"github.com/moricho/tparallel/pkg/ssainstr"
	"golang.org/x/tools/go/ssa"
)

// IsDeferCalled returns whether the given ssa.Function calls `defer`
func IsDeferCalled(f *ssa.Function) bool {
	for _, block := range f.Blocks {
		for _, instr := range block.Instrs {
			switch instr.(type) {
			case *ssa.Defer:
				return true
			}
		}
	}
	return false
}

// IsCalled returns whether the given ssa.Function calls `fn` func
func IsCalled(f *ssa.Function, fn *types.Func) bool {
	block := f.Blocks[0]
	for _, instr := range block.Instrs {
		called := analysisutil.Called(instr, nil, fn)
		if _, ok := ssainstr.LookupCalled(instr, fn); ok || called {
			return true
		}
	}
	return false
}

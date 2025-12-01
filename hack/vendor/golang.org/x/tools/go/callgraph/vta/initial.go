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

// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vta

import (
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/internal/chautil"
	"golang.org/x/tools/go/ssa"
)

// calleesFunc abstracts call graph in one direction,
// from call sites to callees.
type calleesFunc func(ssa.CallInstruction) []*ssa.Function

// makeCalleesFunc returns an initial call graph for vta as a
// calleesFunc. If c is not nil, returns callees as given by c.
// Otherwise, it returns chautil.LazyCallees over fs.
func makeCalleesFunc(fs map[*ssa.Function]bool, c *callgraph.Graph) calleesFunc {
	if c == nil {
		return chautil.LazyCallees(fs)
	}
	return func(call ssa.CallInstruction) []*ssa.Function {
		node := c.Nodes[call.Parent()]
		if node == nil {
			return nil
		}
		var cs []*ssa.Function
		for _, edge := range node.Out {
			if edge.Site == call {
				cs = append(cs, edge.Callee.Func)
			}
		}
		return cs
	}
}

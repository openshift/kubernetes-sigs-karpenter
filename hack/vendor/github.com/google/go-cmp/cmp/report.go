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

// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp

// defaultReporter implements the reporter interface.
//
// As Equal serially calls the PushStep, Report, and PopStep methods, the
// defaultReporter constructs a tree-based representation of the compared value
// and the result of each comparison (see valueNode).
//
// When the String method is called, the FormatDiff method transforms the
// valueNode tree into a textNode tree, which is a tree-based representation
// of the textual output (see textNode).
//
// Lastly, the textNode.String method produces the final report as a string.
type defaultReporter struct {
	root *valueNode
	curr *valueNode
}

func (r *defaultReporter) PushStep(ps PathStep) {
	r.curr = r.curr.PushStep(ps)
	if r.root == nil {
		r.root = r.curr
	}
}
func (r *defaultReporter) Report(rs Result) {
	r.curr.Report(rs)
}
func (r *defaultReporter) PopStep() {
	r.curr = r.curr.PopStep()
}

// String provides a full report of the differences detected as a structured
// literal in pseudo-Go syntax. String may only be called after the entire tree
// has been traversed.
func (r *defaultReporter) String() string {
	assert(r.root != nil && r.curr == nil)
	if r.root.NumDiff == 0 {
		return ""
	}
	ptrs := new(pointerReferences)
	text := formatOptions{}.FormatDiff(r.root, ptrs)
	resolveReferences(text)
	return text.String()
}

func assert(ok bool) {
	if !ok {
		panic("assertion failure")
	}
}

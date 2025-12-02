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

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package nilness inspects the control-flow graph of an SSA function
// and reports errors such as nil pointer dereferences and degenerate
// nil pointer comparisons.
//
// # Analyzer nilness
//
// nilness: check for redundant or impossible nil comparisons
//
// The nilness checker inspects the control-flow graph of each function in
// a package and reports nil pointer dereferences, degenerate nil
// pointers, and panics with nil values. A degenerate comparison is of the form
// x==nil or x!=nil where x is statically known to be nil or non-nil. These are
// often a mistake, especially in control flow related to errors. Panics with nil
// values are checked because they are not detectable by
//
//	if r := recover(); r != nil {
//
// This check reports conditions such as:
//
//	if f == nil { // impossible condition (f is a function)
//	}
//
// and:
//
//	p := &v
//	...
//	if p != nil { // tautological condition
//	}
//
// and:
//
//	if p == nil {
//		print(*p) // nil dereference
//	}
//
// and:
//
//	if p == nil {
//		panic(p)
//	}
//
// Sometimes the control flow may be quite complex, making bugs hard
// to spot. In the example below, the err.Error expression is
// guaranteed to panic because, after the first return, err must be
// nil. The intervening loop is just a distraction.
//
//	...
//	err := g.Wait()
//	if err != nil {
//		return err
//	}
//	partialSuccess := false
//	for _, err := range errs {
//		if err == nil {
//			partialSuccess = true
//			break
//		}
//	}
//	if partialSuccess {
//		reportStatus(StatusMessage{
//			Code:   code.ERROR,
//			Detail: err.Error(), // "nil dereference in dynamic method call"
//		})
//		return nil
//	}
//
// ...
package nilness

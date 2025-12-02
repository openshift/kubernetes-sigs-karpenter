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

// Package unusedwrite checks for unused writes to the elements of a struct or array object.
//
// # Analyzer unusedwrite
//
// unusedwrite: checks for unused writes
//
// The analyzer reports instances of writes to struct fields and
// arrays that are never read. Specifically, when a struct object
// or an array is copied, its elements are copied implicitly by
// the compiler, and any element write to this copy does nothing
// with the original object.
//
// For example:
//
//	type T struct { x int }
//
//	func f(input []T) {
//		for i, v := range input {  // v is a copy
//			v.x = i  // unused write to field x
//		}
//	}
//
// Another example is about non-pointer receiver:
//
//	type T struct { x int }
//
//	func (t T) f() {  // t is a copy
//		t.x = i  // unused write to field x
//	}
package unusedwrite

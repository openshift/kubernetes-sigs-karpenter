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

// Package shadow defines an Analyzer that checks for shadowed variables.
//
// # Analyzer shadow
//
// shadow: check for possible unintended shadowing of variables
//
// This analyzer check for shadowed variables.
// A shadowed variable is a variable declared in an inner scope
// with the same name and type as a variable in an outer scope,
// and where the outer variable is mentioned after the inner one
// is declared.
//
// (This definition can be refined; the module generates too many
// false positives and is not yet enabled by default.)
//
// For example:
//
//	func BadRead(f *os.File, buf []byte) error {
//		var err error
//		for {
//			n, err := f.Read(buf) // shadows the function variable 'err'
//			if err != nil {
//				break // causes return of wrong value
//			}
//			foo(buf)
//		}
//		return err
//	}
package shadow

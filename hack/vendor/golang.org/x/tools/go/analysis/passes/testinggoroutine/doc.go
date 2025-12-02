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

// Package testinggoroutine defines an Analyzerfor detecting calls to
// Fatal from a test goroutine.
//
// # Analyzer testinggoroutine
//
// testinggoroutine: report calls to (*testing.T).Fatal from goroutines started by a test
//
// Functions that abruptly terminate a test, such as the Fatal, Fatalf, FailNow, and
// Skip{,f,Now} methods of *testing.T, must be called from the test goroutine itself.
// This checker detects calls to these functions that occur within a goroutine
// started by the test. For example:
//
//	func TestFoo(t *testing.T) {
//	    go func() {
//	        t.Fatal("oops") // error: (*T).Fatal called from non-test goroutine
//	    }()
//	}
package testinggoroutine

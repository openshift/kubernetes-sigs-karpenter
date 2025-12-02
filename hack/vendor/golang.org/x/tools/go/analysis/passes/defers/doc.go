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

// Package defers defines an Analyzer that checks for common mistakes in defer
// statements.
//
// # Analyzer defers
//
// defers: report common mistakes in defer statements
//
// The defers analyzer reports a diagnostic when a defer statement would
// result in a non-deferred call to time.Since, as experience has shown
// that this is nearly always a mistake.
//
// For example:
//
//	start := time.Now()
//	...
//	defer recordLatency(time.Since(start)) // error: call to time.Since is not deferred
//
// The correct code is:
//
//	defer func() { recordLatency(time.Since(start)) }()
package defers

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

// Package waitgroup defines an Analyzer that detects simple misuses
// of sync.WaitGroup.
//
// # Analyzer waitgroup
//
// waitgroup: check for misuses of sync.WaitGroup
//
// This analyzer detects mistaken calls to the (*sync.WaitGroup).Add
// method from inside a new goroutine, causing Add to race with Wait:
//
//	// WRONG
//	var wg sync.WaitGroup
//	go func() {
//	        wg.Add(1) // "WaitGroup.Add called from inside new goroutine"
//	        defer wg.Done()
//	        ...
//	}()
//	wg.Wait() // (may return prematurely before new goroutine starts)
//
// The correct code calls Add before starting the goroutine:
//
//	// RIGHT
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		...
//	}()
//	wg.Wait()
package waitgroup

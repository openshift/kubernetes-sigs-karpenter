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

// Package unusedresult defines an analyzer that checks for unused
// results of calls to certain pure functions.
//
// # Analyzer unusedresult
//
// unusedresult: check for unused results of calls to some functions
//
// Some functions like fmt.Errorf return a result and have no side
// effects, so it is always a mistake to discard the result. Other
// functions may return an error that must not be ignored, or a cleanup
// operation that must be called. This analyzer reports calls to
// functions like these when the result of the call is ignored.
//
// The set of functions may be controlled using flags.
package unusedresult

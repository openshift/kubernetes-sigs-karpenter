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

// Package appends defines an Analyzer that detects
// if there is only one variable in append.
//
// # Analyzer appends
//
// appends: check for missing values after append
//
// This checker reports calls to append that pass
// no values to be appended to the slice.
//
//	s := []string{"a", "b", "c"}
//	_ = append(s)
//
// Such calls are always no-ops and often indicate an
// underlying mistake.
package appends

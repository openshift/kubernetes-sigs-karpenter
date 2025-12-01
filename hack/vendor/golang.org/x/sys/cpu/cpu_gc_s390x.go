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

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gc

package cpu

// haveAsmFunctions reports whether the other functions in this file can
// be safely called.
func haveAsmFunctions() bool { return true }

// The following feature detection functions are defined in cpu_s390x.s.
// They are likely to be expensive to call so the results should be cached.
func stfle() facilityList
func kmQuery() queryResult
func kmcQuery() queryResult
func kmctrQuery() queryResult
func kmaQuery() queryResult
func kimdQuery() queryResult
func klmdQuery() queryResult

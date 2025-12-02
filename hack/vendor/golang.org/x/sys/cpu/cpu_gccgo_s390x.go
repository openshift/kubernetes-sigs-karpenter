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

//go:build gccgo

package cpu

// haveAsmFunctions reports whether the other functions in this file can
// be safely called.
func haveAsmFunctions() bool { return false }

// TODO(mundaym): the following feature detection functions are currently
// stubs. See https://golang.org/cl/162887 for how to fix this.
// They are likely to be expensive to call so the results should be cached.
func stfle() facilityList     { panic("not implemented for gccgo") }
func kmQuery() queryResult    { panic("not implemented for gccgo") }
func kmcQuery() queryResult   { panic("not implemented for gccgo") }
func kmctrQuery() queryResult { panic("not implemented for gccgo") }
func kmaQuery() queryResult   { panic("not implemented for gccgo") }
func kimdQuery() queryResult  { panic("not implemented for gccgo") }
func klmdQuery() queryResult  { panic("not implemented for gccgo") }

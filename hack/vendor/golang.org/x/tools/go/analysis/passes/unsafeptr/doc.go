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

// Package unsafeptr defines an Analyzer that checks for invalid
// conversions of uintptr to unsafe.Pointer.
//
// # Analyzer unsafeptr
//
// unsafeptr: check for invalid conversions of uintptr to unsafe.Pointer
//
// The unsafeptr analyzer reports likely incorrect uses of unsafe.Pointer
// to convert integers to pointers. A conversion from uintptr to
// unsafe.Pointer is invalid if it implies that there is a uintptr-typed
// word in memory that holds a pointer value, because that word will be
// invisible to stack copying and to the garbage collector.
package unsafeptr

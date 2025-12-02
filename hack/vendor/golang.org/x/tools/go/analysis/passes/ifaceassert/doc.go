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

// Package ifaceassert defines an Analyzer that flags
// impossible interface-interface type assertions.
//
// # Analyzer ifaceassert
//
// ifaceassert: detect impossible interface-to-interface type assertions
//
// This checker flags type assertions v.(T) and corresponding type-switch cases
// in which the static type V of v is an interface that cannot possibly implement
// the target interface T. This occurs when V and T contain methods with the same
// name but different signatures. Example:
//
//	var v interface {
//		Read()
//	}
//	_ = v.(io.Reader)
//
// The Read method in v has a different signature than the Read method in
// io.Reader, so this assertion cannot succeed.
package ifaceassert

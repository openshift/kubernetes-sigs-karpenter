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

// Package reflectvaluecompare defines an Analyzer that checks for accidentally
// using == or reflect.DeepEqual to compare reflect.Value values.
// See issues 43993 and 18871.
//
// # Analyzer reflectvaluecompare
//
// reflectvaluecompare: check for comparing reflect.Value values with == or reflect.DeepEqual
//
// The reflectvaluecompare checker looks for expressions of the form:
//
//	v1 == v2
//	v1 != v2
//	reflect.DeepEqual(v1, v2)
//
// where v1 or v2 are reflect.Values. Comparing reflect.Values directly
// is almost certainly not correct, as it compares the reflect package's
// internal representation, not the underlying value.
// Likely what is intended is:
//
//	v1.Interface() == v2.Interface()
//	v1.Interface() != v2.Interface()
//	reflect.DeepEqual(v1.Interface(), v2.Interface())
package reflectvaluecompare

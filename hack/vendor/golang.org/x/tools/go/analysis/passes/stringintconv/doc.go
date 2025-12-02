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

// Package stringintconv defines an Analyzer that flags type conversions
// from integers to strings.
//
// # Analyzer stringintconv
//
// stringintconv: check for string(int) conversions
//
// This checker flags conversions of the form string(x) where x is an integer
// (but not byte or rune) type. Such conversions are discouraged because they
// return the UTF-8 representation of the Unicode code point x, and not a decimal
// string representation of x as one might expect. Furthermore, if x denotes an
// invalid code point, the conversion cannot be statically rejected.
//
// For conversions that intend on using the code point, consider replacing them
// with string(rune(x)). Otherwise, strconv.Itoa and its equivalents return the
// string representation of the value in the desired base.
package stringintconv

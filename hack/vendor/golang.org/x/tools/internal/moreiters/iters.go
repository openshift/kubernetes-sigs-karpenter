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

// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package moreiters

import "iter"

// First returns the first value of seq and true.
// If seq is empty, it returns the zero value of T and false.
func First[T any](seq iter.Seq[T]) (z T, ok bool) {
	for t := range seq {
		return t, true
	}
	return z, false
}

// Contains reports whether x is an element of the sequence seq.
func Contains[T comparable](seq iter.Seq[T], x T) bool {
	for cand := range seq {
		if cand == x {
			return true
		}
	}
	return false
}

// Every reports whether every pred(t) for t in seq returns true,
// stopping at the first false element.
func Every[T any](seq iter.Seq[T], pred func(T) bool) bool {
	for t := range seq {
		if !pred(t) {
			return false
		}
	}
	return true
}

// Any reports whether any pred(t) for t in seq returns true.
func Any[T any](seq iter.Seq[T], pred func(T) bool) bool {
	for t := range seq {
		if pred(t) {
			return true
		}
	}
	return false
}

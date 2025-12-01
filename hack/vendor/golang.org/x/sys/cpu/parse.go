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

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

import "strconv"

// parseRelease parses a dot-separated version number. It follows the semver
// syntax, but allows the minor and patch versions to be elided.
//
// This is a copy of the Go runtime's parseRelease from
// https://golang.org/cl/209597.
func parseRelease(rel string) (major, minor, patch int, ok bool) {
	// Strip anything after a dash or plus.
	for i := range len(rel) {
		if rel[i] == '-' || rel[i] == '+' {
			rel = rel[:i]
			break
		}
	}

	next := func() (int, bool) {
		for i := range len(rel) {
			if rel[i] == '.' {
				ver, err := strconv.Atoi(rel[:i])
				rel = rel[i+1:]
				return ver, err == nil
			}
		}
		ver, err := strconv.Atoi(rel)
		rel = ""
		return ver, err == nil
	}
	if major, ok = next(); !ok || rel == "" {
		return
	}
	if minor, ok = next(); !ok || rel == "" {
		return
	}
	patch, ok = next()
	return
}

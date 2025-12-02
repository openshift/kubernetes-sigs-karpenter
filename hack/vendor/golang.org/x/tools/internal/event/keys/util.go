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

package keys

import (
	"sort"
	"strings"
)

// Join returns a canonical join of the keys in S:
// a sorted comma-separated string list.
func Join[S ~[]T, T ~string](s S) string {
	strs := make([]string, 0, len(s))
	for _, v := range s {
		strs = append(strs, string(v))
	}
	sort.Strings(strs)
	return strings.Join(strs, ",")
}

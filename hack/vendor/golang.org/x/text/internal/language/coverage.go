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

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package language

// BaseLanguages returns the list of all supported base languages. It generates
// the list by traversing the internal structures.
func BaseLanguages() []Language {
	base := make([]Language, 0, NumLanguages)
	for i := 0; i < langNoIndexOffset; i++ {
		// We included "und" already for the value 0.
		if i != nonCanonicalUnd {
			base = append(base, Language(i))
		}
	}
	i := langNoIndexOffset
	for _, v := range langNoIndex {
		for k := 0; k < 8; k++ {
			if v&1 == 1 {
				base = append(base, Language(i))
			}
			v >>= 1
			i++
		}
	}
	return base
}

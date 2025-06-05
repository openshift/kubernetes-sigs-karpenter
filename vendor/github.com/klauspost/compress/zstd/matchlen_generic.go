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

//go:build !amd64 || appengine || !gc || noasm
// +build !amd64 appengine !gc noasm

// Copyright 2019+ Klaus Post. All rights reserved.
// License information can be found in the LICENSE file.

package zstd

import (
	"encoding/binary"
	"math/bits"
)

// matchLen returns the maximum common prefix length of a and b.
// a must be the shortest of the two.
func matchLen(a, b []byte) (n int) {
	for ; len(a) >= 8 && len(b) >= 8; a, b = a[8:], b[8:] {
		diff := binary.LittleEndian.Uint64(a) ^ binary.LittleEndian.Uint64(b)
		if diff != 0 {
			return n + bits.TrailingZeros64(diff)>>3
		}
		n += 8
	}

	for i := range a {
		if a[i] != b[i] {
			break
		}
		n++
	}
	return n

}

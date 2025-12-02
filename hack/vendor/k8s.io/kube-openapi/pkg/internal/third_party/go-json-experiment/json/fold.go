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

// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"unicode"
	"unicode/utf8"
)

// foldName returns a folded string such that foldName(x) == foldName(y)
// is similar to strings.EqualFold(x, y), but ignores underscore and dashes.
// This allows foldName to match common naming conventions.
func foldName(in []byte) []byte {
	// This is inlinable to take advantage of "function outlining".
	// See https://blog.filippo.io/efficient-go-apis-with-the-inliner/
	var arr [32]byte // large enough for most JSON names
	return appendFoldedName(arr[:0], in)
}
func appendFoldedName(out, in []byte) []byte {
	for i := 0; i < len(in); {
		// Handle single-byte ASCII.
		if c := in[i]; c < utf8.RuneSelf {
			if c != '_' && c != '-' {
				if 'a' <= c && c <= 'z' {
					c -= 'a' - 'A'
				}
				out = append(out, c)
			}
			i++
			continue
		}
		// Handle multi-byte Unicode.
		r, n := utf8.DecodeRune(in[i:])
		out = utf8.AppendRune(out, foldRune(r))
		i += n
	}
	return out
}

// foldRune is a variation on unicode.SimpleFold that returns the same rune
// for all runes in the same fold set.
//
// Invariant:
//
//	foldRune(x) == foldRune(y) ⇔ strings.EqualFold(string(x), string(y))
func foldRune(r rune) rune {
	for {
		r2 := unicode.SimpleFold(r)
		if r2 <= r {
			return r2 // smallest character in the fold set
		}
		r = r2
	}
}

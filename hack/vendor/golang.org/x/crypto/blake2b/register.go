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

// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blake2b

import (
	"crypto"
	"hash"
)

func init() {
	newHash256 := func() hash.Hash {
		h, _ := New256(nil)
		return h
	}
	newHash384 := func() hash.Hash {
		h, _ := New384(nil)
		return h
	}

	newHash512 := func() hash.Hash {
		h, _ := New512(nil)
		return h
	}

	crypto.RegisterHash(crypto.BLAKE2b_256, newHash256)
	crypto.RegisterHash(crypto.BLAKE2b_384, newHash384)
	crypto.RegisterHash(crypto.BLAKE2b_512, newHash512)
}

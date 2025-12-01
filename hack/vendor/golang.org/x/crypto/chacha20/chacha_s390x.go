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

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gc && !purego

package chacha20

import "golang.org/x/sys/cpu"

var haveAsm = cpu.S390X.HasVX

const bufSize = 256

// xorKeyStreamVX is an assembly implementation of XORKeyStream. It must only
// be called when the vector facility is available. Implementation in asm_s390x.s.
//
//go:noescape
func xorKeyStreamVX(dst, src []byte, key *[8]uint32, nonce *[3]uint32, counter *uint32)

func (c *Cipher) xorKeyStreamBlocks(dst, src []byte) {
	if cpu.S390X.HasVX {
		xorKeyStreamVX(dst, src, &c.key, &c.nonce, &c.counter)
	} else {
		c.xorKeyStreamBlocksGeneric(dst, src)
	}
}

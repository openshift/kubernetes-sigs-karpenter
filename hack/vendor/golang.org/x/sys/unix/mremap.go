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

//go:build linux || netbsd

package unix

import "unsafe"

type mremapMmapper struct {
	mmapper
	mremap func(oldaddr uintptr, oldlength uintptr, newlength uintptr, flags int, newaddr uintptr) (xaddr uintptr, err error)
}

var mapper = &mremapMmapper{
	mmapper: mmapper{
		active: make(map[*byte][]byte),
		mmap:   mmap,
		munmap: munmap,
	},
	mremap: mremap,
}

func (m *mremapMmapper) Mremap(oldData []byte, newLength int, flags int) (data []byte, err error) {
	if newLength <= 0 || len(oldData) == 0 || len(oldData) != cap(oldData) || flags&mremapFixed != 0 {
		return nil, EINVAL
	}

	pOld := &oldData[cap(oldData)-1]
	m.Lock()
	defer m.Unlock()
	bOld := m.active[pOld]
	if bOld == nil || &bOld[0] != &oldData[0] {
		return nil, EINVAL
	}
	newAddr, errno := m.mremap(uintptr(unsafe.Pointer(&bOld[0])), uintptr(len(bOld)), uintptr(newLength), flags, 0)
	if errno != nil {
		return nil, errno
	}
	bNew := unsafe.Slice((*byte)(unsafe.Pointer(newAddr)), newLength)
	pNew := &bNew[cap(bNew)-1]
	if flags&mremapDontunmap == 0 {
		delete(m.active, pOld)
	}
	m.active[pNew] = bNew
	return bNew, nil
}

func Mremap(oldData []byte, newLength int, flags int) (data []byte, err error) {
	return mapper.Mremap(oldData, newLength, flags)
}

func MremapPtr(oldAddr unsafe.Pointer, oldSize uintptr, newAddr unsafe.Pointer, newSize uintptr, flags int) (ret unsafe.Pointer, err error) {
	xaddr, err := mapper.mremap(uintptr(oldAddr), oldSize, newSize, flags, uintptr(newAddr))
	return unsafe.Pointer(xaddr), err
}

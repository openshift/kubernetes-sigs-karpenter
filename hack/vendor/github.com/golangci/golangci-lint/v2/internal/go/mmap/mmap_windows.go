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

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mmap

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func mmapFile(f *os.File) (Data, error) {
	st, err := f.Stat()
	if err != nil {
		return Data{}, err
	}
	size := st.Size()
	if size == 0 {
		return Data{f, nil}, nil
	}
	h, err := syscall.CreateFileMapping(syscall.Handle(f.Fd()), nil, syscall.PAGE_READONLY, 0, 0, nil)
	if err != nil {
		return Data{}, fmt.Errorf("CreateFileMapping %s: %w", f.Name(), err)
	}

	addr, err := syscall.MapViewOfFile(h, syscall.FILE_MAP_READ, 0, 0, 0)
	if err != nil {
		return Data{}, fmt.Errorf("MapViewOfFile %s: %w", f.Name(), err)
	}
	var info windows.MemoryBasicInformation
	err = windows.VirtualQuery(addr, &info, unsafe.Sizeof(info))
	if err != nil {
		return Data{}, fmt.Errorf("VirtualQuery %s: %w", f.Name(), err)
	}
	data := unsafe.Slice((*byte)(unsafe.Pointer(addr)), int(info.RegionSize))
	if len(data) < int(size) {
		// In some cases, especially on 386, we may not receive a in incomplete mapping:
		// one that is shorter than the file itself. Return an error in those cases because
		// incomplete mappings are not useful.
		return Data{}, fmt.Errorf("mmapFile: received incomplete mapping of file")
	}
	return Data{f, data[:int(size)]}, nil
}

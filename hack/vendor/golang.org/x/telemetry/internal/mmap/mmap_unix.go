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

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package mmap

import (
	"fmt"
	"io/fs"
	"os"
	"syscall"
)

func mmapFile(f *os.File) (*Data, error) {
	st, err := f.Stat()
	if err != nil {
		return nil, err
	}
	size := st.Size()
	pagesize := int64(os.Getpagesize())
	if int64(int(size+(pagesize-1))) != size+(pagesize-1) {
		return nil, fmt.Errorf("%s: too large for mmap", f.Name())
	}
	n := int(size)
	if n == 0 {
		return &Data{f, nil, nil}, nil
	}
	mmapLength := int(((size + pagesize - 1) / pagesize) * pagesize) // round up to page size
	data, err := syscall.Mmap(int(f.Fd()), 0, mmapLength, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return nil, &fs.PathError{Op: "mmap", Path: f.Name(), Err: err}
	}
	return &Data{f, data[:n], nil}, nil
}

func munmapFile(d *Data) error {
	if len(d.Data) == 0 {
		return nil
	}
	err := syscall.Munmap(d.Data)
	if err != nil {
		return &fs.PathError{Op: "munmap", Path: d.f.Name(), Err: err}
	}
	return nil
}

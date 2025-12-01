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

// This package is a lightly modified version of the mmap code
// in github.com/google/codesearch/index.

// The mmap package provides an abstraction for memory mapping files
// on different platforms.
package mmap

import (
	"os"
)

// Data is mmap'ed read-only data from a file.
// The backing file is never closed, so Data
// remains valid for the lifetime of the process.
type Data struct {
	f    *os.File
	Data []byte
}

// Mmap maps the given file into memory.
func Mmap(file string) (Data, bool, error) {
	f, err := os.Open(file)
	if err != nil {
		return Data{}, false, err
	}
	data, err := mmapFile(f)
	return data, true, err
}

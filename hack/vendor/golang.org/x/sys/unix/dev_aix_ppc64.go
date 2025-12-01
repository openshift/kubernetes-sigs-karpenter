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

//go:build aix && ppc64

// Functions to access/create device major and minor numbers matching the
// encoding used AIX.

package unix

// Major returns the major component of a Linux device number.
func Major(dev uint64) uint32 {
	return uint32((dev & 0x3fffffff00000000) >> 32)
}

// Minor returns the minor component of a Linux device number.
func Minor(dev uint64) uint32 {
	return uint32((dev & 0x00000000ffffffff) >> 0)
}

// Mkdev returns a Linux device number generated from the given major and minor
// components.
func Mkdev(major, minor uint32) uint64 {
	var DEVNO64 uint64
	DEVNO64 = 0x8000000000000000
	return ((uint64(major) << 32) | (uint64(minor) & 0x00000000FFFFFFFF) | DEVNO64)
}

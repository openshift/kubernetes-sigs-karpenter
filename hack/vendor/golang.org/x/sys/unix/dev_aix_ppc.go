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

//go:build aix && ppc

// Functions to access/create device major and minor numbers matching the
// encoding used by AIX.

package unix

// Major returns the major component of a Linux device number.
func Major(dev uint64) uint32 {
	return uint32((dev >> 16) & 0xffff)
}

// Minor returns the minor component of a Linux device number.
func Minor(dev uint64) uint32 {
	return uint32(dev & 0xffff)
}

// Mkdev returns a Linux device number generated from the given major and minor
// components.
func Mkdev(major, minor uint32) uint64 {
	return uint64(((major) << 16) | (minor))
}

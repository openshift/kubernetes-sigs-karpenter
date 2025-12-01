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

// Functions to access/create device major and minor numbers matching the
// encoding used in Darwin's sys/types.h header.

package unix

// Major returns the major component of a Darwin device number.
func Major(dev uint64) uint32 {
	return uint32((dev >> 24) & 0xff)
}

// Minor returns the minor component of a Darwin device number.
func Minor(dev uint64) uint32 {
	return uint32(dev & 0xffffff)
}

// Mkdev returns a Darwin device number generated from the given major and minor
// components.
func Mkdev(major, minor uint32) uint64 {
	return (uint64(major) << 24) | uint64(minor)
}

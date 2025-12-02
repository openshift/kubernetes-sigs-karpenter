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
// encoding used in FreeBSD's sys/types.h header.
//
// The information below is extracted and adapted from sys/types.h:
//
// Minor gives a cookie instead of an index since in order to avoid changing the
// meanings of bits 0-15 or wasting time and space shifting bits 16-31 for
// devices that don't use them.

package unix

// Major returns the major component of a FreeBSD device number.
func Major(dev uint64) uint32 {
	return uint32((dev >> 8) & 0xff)
}

// Minor returns the minor component of a FreeBSD device number.
func Minor(dev uint64) uint32 {
	return uint32(dev & 0xffff00ff)
}

// Mkdev returns a FreeBSD device number generated from the given major and
// minor components.
func Mkdev(major, minor uint32) uint64 {
	return (uint64(major) << 8) | uint64(minor)
}

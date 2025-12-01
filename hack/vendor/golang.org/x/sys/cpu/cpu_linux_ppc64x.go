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

//go:build linux && (ppc64 || ppc64le)

package cpu

// HWCAP/HWCAP2 bits. These are exposed by the kernel.
const (
	// ISA Level
	_PPC_FEATURE2_ARCH_2_07 = 0x80000000
	_PPC_FEATURE2_ARCH_3_00 = 0x00800000

	// CPU features
	_PPC_FEATURE2_DARN = 0x00200000
	_PPC_FEATURE2_SCV  = 0x00100000
)

func doinit() {
	// HWCAP2 feature bits
	PPC64.IsPOWER8 = isSet(hwCap2, _PPC_FEATURE2_ARCH_2_07)
	PPC64.IsPOWER9 = isSet(hwCap2, _PPC_FEATURE2_ARCH_3_00)
	PPC64.HasDARN = isSet(hwCap2, _PPC_FEATURE2_DARN)
	PPC64.HasSCV = isSet(hwCap2, _PPC_FEATURE2_SCV)
}

func isSet(hwc uint, value uint) bool {
	return hwc&value != 0
}

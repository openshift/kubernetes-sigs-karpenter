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

// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

// HWCAP bits. These are exposed by the Linux kernel.
const (
	hwcap_LOONGARCH_LSX  = 1 << 4
	hwcap_LOONGARCH_LASX = 1 << 5
)

func doinit() {
	// TODO: Features that require kernel support like LSX and LASX can
	// be detected here once needed in std library or by the compiler.
	Loong64.HasLSX = hwcIsSet(hwCap, hwcap_LOONGARCH_LSX)
	Loong64.HasLASX = hwcIsSet(hwCap, hwcap_LOONGARCH_LASX)
}

func hwcIsSet(hwc uint, val uint) bool {
	return hwc&val != 0
}

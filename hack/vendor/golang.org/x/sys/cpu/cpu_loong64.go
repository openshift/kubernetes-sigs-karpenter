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

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong64

package cpu

const cacheLineSize = 64

// Bit fields for CPUCFG registers, Related reference documents:
// https://loongson.github.io/LoongArch-Documentation/LoongArch-Vol1-EN.html#_cpucfg
const (
	// CPUCFG1 bits
	cpucfg1_CRC32 = 1 << 25

	// CPUCFG2 bits
	cpucfg2_LAM_BH = 1 << 27
	cpucfg2_LAMCAS = 1 << 28
)

func initOptions() {
	options = []option{
		{Name: "lsx", Feature: &Loong64.HasLSX},
		{Name: "lasx", Feature: &Loong64.HasLASX},
		{Name: "crc32", Feature: &Loong64.HasCRC32},
		{Name: "lam_bh", Feature: &Loong64.HasLAM_BH},
		{Name: "lamcas", Feature: &Loong64.HasLAMCAS},
	}

	// The CPUCFG data on Loong64 only reflects the hardware capabilities,
	// not the kernel support status, so features such as LSX and LASX that
	// require kernel support cannot be obtained from the CPUCFG data.
	//
	// These features only require hardware capability support and do not
	// require kernel specific support, so they can be obtained directly
	// through CPUCFG
	cfg1 := get_cpucfg(1)
	cfg2 := get_cpucfg(2)

	Loong64.HasCRC32 = cfgIsSet(cfg1, cpucfg1_CRC32)
	Loong64.HasLAMCAS = cfgIsSet(cfg2, cpucfg2_LAMCAS)
	Loong64.HasLAM_BH = cfgIsSet(cfg2, cpucfg2_LAM_BH)
}

func get_cpucfg(reg uint32) uint32

func cfgIsSet(cfg uint32, val uint32) bool {
	return cfg&val != 0
}

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

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix

package cpu

const (
	// getsystemcfg constants
	_SC_IMPL     = 2
	_IMPL_POWER8 = 0x10000
	_IMPL_POWER9 = 0x20000
)

func archInit() {
	impl := getsystemcfg(_SC_IMPL)
	if impl&_IMPL_POWER8 != 0 {
		PPC64.IsPOWER8 = true
	}
	if impl&_IMPL_POWER9 != 0 {
		PPC64.IsPOWER8 = true
		PPC64.IsPOWER9 = true
	}

	Initialized = true
}

func getsystemcfg(label int) (n uint64) {
	r0, _ := callgetsystemcfg(label)
	n = uint64(r0)
	return
}

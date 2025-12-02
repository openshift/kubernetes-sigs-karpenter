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

// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Recreate a getsystemcfg syscall handler instead of
// using the one provided by x/sys/unix to avoid having
// the dependency between them. (See golang.org/issue/32102)
// Moreover, this file will be used during the building of
// gccgo's libgo and thus must not used a CGo method.

//go:build aix && gccgo

package cpu

import (
	"syscall"
)

//extern getsystemcfg
func gccgoGetsystemcfg(label uint32) (r uint64)

func callgetsystemcfg(label int) (r1 uintptr, e1 syscall.Errno) {
	r1 = uintptr(gccgoGetsystemcfg(uint32(label)))
	e1 = syscall.GetErrno()
	return
}

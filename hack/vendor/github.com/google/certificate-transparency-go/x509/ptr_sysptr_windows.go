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

//go:build go1.11
// +build go1.11

package x509

import (
	"syscall"
	"unsafe"
)

// For Go versions >= 1.11, the ExtraPolicyPara field in
// syscall.CertChainPolicyPara is of type syscall.Pointer.  See:
//   https://github.com/golang/go/commit/4869ec00e87ef

func convertToPolicyParaType(p unsafe.Pointer) syscall.Pointer {
	return (syscall.Pointer)(p)
}

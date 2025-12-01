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

// Copyright 2018, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package value

import (
	"reflect"
	"unsafe"
)

// Pointer is an opaque typed pointer and is guaranteed to be comparable.
type Pointer struct {
	p unsafe.Pointer
	t reflect.Type
}

// PointerOf returns a Pointer from v, which must be a
// reflect.Ptr, reflect.Slice, or reflect.Map.
func PointerOf(v reflect.Value) Pointer {
	// The proper representation of a pointer is unsafe.Pointer,
	// which is necessary if the GC ever uses a moving collector.
	return Pointer{unsafe.Pointer(v.Pointer()), v.Type()}
}

// IsNil reports whether the pointer is nil.
func (p Pointer) IsNil() bool {
	return p.p == nil
}

// Uintptr returns the pointer as a uintptr.
func (p Pointer) Uintptr() uintptr {
	return uintptr(p.p)
}

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

//+build !go1.18

package reflect2

import (
	"unsafe"
)

// m escapes into the return value, but the caller of mapiterinit
// doesn't let the return value escape.
//go:noescape
//go:linkname mapiterinit reflect.mapiterinit
func mapiterinit(rtype unsafe.Pointer, m unsafe.Pointer) (val *hiter)

func (type2 *UnsafeMapType) UnsafeIterate(obj unsafe.Pointer) MapIterator {
	return &UnsafeMapIterator{
		hiter:      mapiterinit(type2.rtype, *(*unsafe.Pointer)(obj)),
		pKeyRType:  type2.pKeyRType,
		pElemRType: type2.pElemRType,
	}
}
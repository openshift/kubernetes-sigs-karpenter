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

package danger

import (
	"reflect"
	"unsafe"
)

// typeID is used as key in encoder and decoder caches to enable using
// the optimize runtime.mapaccess2_fast64 function instead of the more
// expensive lookup if we were to use reflect.Type as map key.
//
// typeID holds the pointer to the reflect.Type value, which is unique
// in the program.
//
// https://github.com/segmentio/encoding/blob/master/json/codec.go#L59-L61
type TypeID unsafe.Pointer

func MakeTypeID(t reflect.Type) TypeID {
	// reflect.Type has the fields:
	// typ unsafe.Pointer
	// ptr unsafe.Pointer
	return TypeID((*[2]unsafe.Pointer)(unsafe.Pointer(&t))[1])
}

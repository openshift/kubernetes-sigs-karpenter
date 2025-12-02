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

package conv

import "unsafe"

// bitsize returns the size in bits of a type.
//
// NOTE: [unsafe.SizeOf] simply returns the size in bytes of the value.
// For primitive types T, the generic stencil is precompiled and this value
// is resolved at compile time, resulting in an immediate call to [strconv.ParseFloat].
//
// We may leave up to the go compiler to simplify this function into a
// constant value, which happens in practice at least for primitive types
// (e.g. numerical types).
func bitsize[T Numerical](value T) int {
	const bitsPerByte = 8
	return int(unsafe.Sizeof(value)) * bitsPerByte
}

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

package types

// AsArray is a type-assert like operation, x.(*Array), but never panics.
// Returns nil if type is not an array.
func AsArray(x Type) *Array { return nil }

// AsSlice is a type-assert like operation, x.(*Slice), but never panics.
// Returns nil if type is not an array.
func AsSlice(x Type) *Slice { return nil }

// AsPointer is a type-assert like operation, x.(*Pointer), but never panics.
// Returns nil if type is not a pointer.
func AsPointer(x Type) *Pointer { return nil }

// AsStruct is a type-assert like operation, x.(*Struct), but never panics.
// Returns nil if type is not a struct.
func AsStruct(x Type) *Struct { return nil }

// AsInterface is a type-assert like operation, x.(*Interface), but never panics.
// Returns nil if type is not an interface.
func AsInterface(x Type) *Interface { return nil }

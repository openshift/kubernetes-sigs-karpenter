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

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astutil

import (
	"go/ast"
	"reflect"
)

// CloneNode returns a deep copy of a Node.
// It omits pointers to ast.{Scope,Object} variables.
func CloneNode[T ast.Node](n T) T {
	return cloneNode(n).(T)
}

func cloneNode(n ast.Node) ast.Node {
	var clone func(x reflect.Value) reflect.Value
	set := func(dst, src reflect.Value) {
		src = clone(src)
		if src.IsValid() {
			dst.Set(src)
		}
	}
	clone = func(x reflect.Value) reflect.Value {
		switch x.Kind() {
		case reflect.Pointer:
			if x.IsNil() {
				return x
			}
			// Skip fields of types potentially involved in cycles.
			switch x.Interface().(type) {
			case *ast.Object, *ast.Scope:
				return reflect.Zero(x.Type())
			}
			y := reflect.New(x.Type().Elem())
			set(y.Elem(), x.Elem())
			return y

		case reflect.Struct:
			y := reflect.New(x.Type()).Elem()
			for i := 0; i < x.Type().NumField(); i++ {
				set(y.Field(i), x.Field(i))
			}
			return y

		case reflect.Slice:
			if x.IsNil() {
				return x
			}
			y := reflect.MakeSlice(x.Type(), x.Len(), x.Cap())
			for i := 0; i < x.Len(); i++ {
				set(y.Index(i), x.Index(i))
			}
			return y

		case reflect.Interface:
			y := reflect.New(x.Type()).Elem()
			set(y, x.Elem())
			return y

		case reflect.Array, reflect.Chan, reflect.Func, reflect.Map, reflect.UnsafePointer:
			panic(x) // unreachable in AST

		default:
			return x // bool, string, number
		}
	}
	return clone(reflect.ValueOf(n)).Interface().(ast.Node)
}

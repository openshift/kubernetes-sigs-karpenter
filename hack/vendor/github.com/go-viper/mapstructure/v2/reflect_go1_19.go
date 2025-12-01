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

//go:build !go1.20

package mapstructure

import "reflect"

func isComparable(v reflect.Value) bool {
	k := v.Kind()
	switch k {
	case reflect.Invalid:
		return false

	case reflect.Array:
		switch v.Type().Elem().Kind() {
		case reflect.Interface, reflect.Array, reflect.Struct:
			for i := 0; i < v.Type().Len(); i++ {
				// if !v.Index(i).Comparable() {
				if !isComparable(v.Index(i)) {
					return false
				}
			}
			return true
		}
		return v.Type().Comparable()

	case reflect.Interface:
		// return v.Elem().Comparable()
		return isComparable(v.Elem())

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			return false

			// if !v.Field(i).Comparable() {
			if !isComparable(v.Field(i)) {
				return false
			}
		}
		return true

	default:
		return v.Type().Comparable()
	}
}

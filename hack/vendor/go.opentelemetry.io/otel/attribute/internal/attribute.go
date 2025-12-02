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

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

/*
Package attribute provide several helper functions for some commonly used
logic of processing attributes.
*/
package attribute // import "go.opentelemetry.io/otel/attribute/internal"

import (
	"reflect"
)

// BoolSliceValue converts a bool slice into an array with same elements as slice.
func BoolSliceValue(v []bool) interface{} {
	var zero bool
	cp := reflect.New(reflect.ArrayOf(len(v), reflect.TypeOf(zero))).Elem()
	reflect.Copy(cp, reflect.ValueOf(v))
	return cp.Interface()
}

// Int64SliceValue converts an int64 slice into an array with same elements as slice.
func Int64SliceValue(v []int64) interface{} {
	var zero int64
	cp := reflect.New(reflect.ArrayOf(len(v), reflect.TypeOf(zero))).Elem()
	reflect.Copy(cp, reflect.ValueOf(v))
	return cp.Interface()
}

// Float64SliceValue converts a float64 slice into an array with same elements as slice.
func Float64SliceValue(v []float64) interface{} {
	var zero float64
	cp := reflect.New(reflect.ArrayOf(len(v), reflect.TypeOf(zero))).Elem()
	reflect.Copy(cp, reflect.ValueOf(v))
	return cp.Interface()
}

// StringSliceValue converts a string slice into an array with same elements as slice.
func StringSliceValue(v []string) interface{} {
	var zero string
	cp := reflect.New(reflect.ArrayOf(len(v), reflect.TypeOf(zero))).Elem()
	reflect.Copy(cp, reflect.ValueOf(v))
	return cp.Interface()
}

// AsBoolSlice converts a bool array into a slice into with same elements as array.
func AsBoolSlice(v interface{}) []bool {
	rv := reflect.ValueOf(v)
	if rv.Type().Kind() != reflect.Array {
		return nil
	}
	cpy := make([]bool, rv.Len())
	if len(cpy) > 0 {
		_ = reflect.Copy(reflect.ValueOf(cpy), rv)
	}
	return cpy
}

// AsInt64Slice converts an int64 array into a slice into with same elements as array.
func AsInt64Slice(v interface{}) []int64 {
	rv := reflect.ValueOf(v)
	if rv.Type().Kind() != reflect.Array {
		return nil
	}
	cpy := make([]int64, rv.Len())
	if len(cpy) > 0 {
		_ = reflect.Copy(reflect.ValueOf(cpy), rv)
	}
	return cpy
}

// AsFloat64Slice converts a float64 array into a slice into with same elements as array.
func AsFloat64Slice(v interface{}) []float64 {
	rv := reflect.ValueOf(v)
	if rv.Type().Kind() != reflect.Array {
		return nil
	}
	cpy := make([]float64, rv.Len())
	if len(cpy) > 0 {
		_ = reflect.Copy(reflect.ValueOf(cpy), rv)
	}
	return cpy
}

// AsStringSlice converts a string array into a slice into with same elements as array.
func AsStringSlice(v interface{}) []string {
	rv := reflect.ValueOf(v)
	if rv.Type().Kind() != reflect.Array {
		return nil
	}
	cpy := make([]string, rv.Len())
	if len(cpy) > 0 {
		_ = reflect.Copy(reflect.ValueOf(cpy), rv)
	}
	return cpy
}

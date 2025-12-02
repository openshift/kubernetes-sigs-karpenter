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

// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package cast provides easy and safe casting in Go.
package cast

import "time"

const errorMsg = "unable to cast %#v of type %T to %T"
const errorMsgWith = "unable to cast %#v of type %T to %T: %w"

// Basic is a type parameter constraint for functions accepting basic types.
//
// It represents the supported basic types this package can cast to.
type Basic interface {
	string | bool | Number | time.Time | time.Duration
}

// ToE casts any value to a [Basic] type.
func ToE[T Basic](i any) (T, error) {
	var t T

	var v any
	var err error

	switch any(t).(type) {
	case string:
		v, err = ToStringE(i)
	case bool:
		v, err = ToBoolE(i)
	case int:
		v, err = toNumberE[int](i, parseInt[int])
	case int8:
		v, err = toNumberE[int8](i, parseInt[int8])
	case int16:
		v, err = toNumberE[int16](i, parseInt[int16])
	case int32:
		v, err = toNumberE[int32](i, parseInt[int32])
	case int64:
		v, err = toNumberE[int64](i, parseInt[int64])
	case uint:
		v, err = toUnsignedNumberE[uint](i, parseUint[uint])
	case uint8:
		v, err = toUnsignedNumberE[uint8](i, parseUint[uint8])
	case uint16:
		v, err = toUnsignedNumberE[uint16](i, parseUint[uint16])
	case uint32:
		v, err = toUnsignedNumberE[uint32](i, parseUint[uint32])
	case uint64:
		v, err = toUnsignedNumberE[uint64](i, parseUint[uint64])
	case float32:
		v, err = toNumberE[float32](i, parseFloat[float32])
	case float64:
		v, err = toNumberE[float64](i, parseFloat[float64])
	case time.Time:
		v, err = ToTimeE(i)
	case time.Duration:
		v, err = ToDurationE(i)
	}

	if err != nil {
		return t, err
	}

	return v.(T), nil
}

// Must is a helper that wraps a call to a cast function and panics if the error is non-nil.
func Must[T any](i any, err error) T {
	if err != nil {
		panic(err)
	}

	return i.(T)
}

// To casts any value to a [Basic] type.
func To[T Basic](i any) T {
	v, _ := ToE[T](i)

	return v
}

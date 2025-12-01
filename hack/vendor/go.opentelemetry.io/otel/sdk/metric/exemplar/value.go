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

package exemplar // import "go.opentelemetry.io/otel/sdk/metric/exemplar"

import "math"

// ValueType identifies the type of value used in exemplar data.
type ValueType uint8

const (
	// UnknownValueType should not be used. It represents a misconfigured
	// Value.
	UnknownValueType ValueType = 0
	// Int64ValueType represents a Value with int64 data.
	Int64ValueType ValueType = 1
	// Float64ValueType represents a Value with float64 data.
	Float64ValueType ValueType = 2
)

// Value is the value of data held by an exemplar.
type Value struct {
	t   ValueType
	val uint64
}

// NewValue returns a new [Value] for the provided value.
func NewValue[N int64 | float64](value N) Value {
	switch v := any(value).(type) {
	case int64:
		// This can be later converted back to int64 (overflow not checked).
		return Value{t: Int64ValueType, val: uint64(v)} // nolint:gosec
	case float64:
		return Value{t: Float64ValueType, val: math.Float64bits(v)}
	}
	return Value{}
}

// Type returns the [ValueType] of data held by v.
func (v Value) Type() ValueType { return v.t }

// Int64 returns the value of v as an int64. If the ValueType of v is not an
// Int64ValueType, 0 is returned.
func (v Value) Int64() int64 {
	if v.t == Int64ValueType {
		// Assumes the correct int64 was stored in v.val based on type.
		return int64(v.val) // nolint: gosec
	}
	return 0
}

// Float64 returns the value of v as an float64. If the ValueType of v is not
// an Float64ValueType, 0 is returned.
func (v Value) Float64() float64 {
	if v.t == Float64ValueType {
		return math.Float64frombits(v.val)
	}
	return 0
}

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

package telemetry // import "go.opentelemetry.io/otel/trace/internal/telemetry"

// Attr is a key-value pair.
type Attr struct {
	Key   string `json:"key,omitempty"`
	Value Value  `json:"value,omitempty"`
}

// String returns an Attr for a string value.
func String(key, value string) Attr {
	return Attr{key, StringValue(value)}
}

// Int64 returns an Attr for an int64 value.
func Int64(key string, value int64) Attr {
	return Attr{key, Int64Value(value)}
}

// Int returns an Attr for an int value.
func Int(key string, value int) Attr {
	return Int64(key, int64(value))
}

// Float64 returns an Attr for a float64 value.
func Float64(key string, value float64) Attr {
	return Attr{key, Float64Value(value)}
}

// Bool returns an Attr for a bool value.
func Bool(key string, value bool) Attr {
	return Attr{key, BoolValue(value)}
}

// Bytes returns an Attr for a []byte value.
// The passed slice must not be changed after it is passed.
func Bytes(key string, value []byte) Attr {
	return Attr{key, BytesValue(value)}
}

// Slice returns an Attr for a []Value value.
// The passed slice must not be changed after it is passed.
func Slice(key string, value ...Value) Attr {
	return Attr{key, SliceValue(value...)}
}

// Map returns an Attr for a map value.
// The passed slice must not be changed after it is passed.
func Map(key string, value ...Attr) Attr {
	return Attr{key, MapValue(value...)}
}

// Equal returns if a is equal to b.
func (a Attr) Equal(b Attr) bool {
	return a.Key == b.Key && a.Value.Equal(b.Value)
}

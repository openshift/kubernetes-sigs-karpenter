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

package jsoniter

import (
	"strconv"
)

type floatAny struct {
	baseAny
	val float64
}

func (any *floatAny) Parse() *Iterator {
	return nil
}

func (any *floatAny) ValueType() ValueType {
	return NumberValue
}

func (any *floatAny) MustBeValid() Any {
	return any
}

func (any *floatAny) LastError() error {
	return nil
}

func (any *floatAny) ToBool() bool {
	return any.ToFloat64() != 0
}

func (any *floatAny) ToInt() int {
	return int(any.val)
}

func (any *floatAny) ToInt32() int32 {
	return int32(any.val)
}

func (any *floatAny) ToInt64() int64 {
	return int64(any.val)
}

func (any *floatAny) ToUint() uint {
	if any.val > 0 {
		return uint(any.val)
	}
	return 0
}

func (any *floatAny) ToUint32() uint32 {
	if any.val > 0 {
		return uint32(any.val)
	}
	return 0
}

func (any *floatAny) ToUint64() uint64 {
	if any.val > 0 {
		return uint64(any.val)
	}
	return 0
}

func (any *floatAny) ToFloat32() float32 {
	return float32(any.val)
}

func (any *floatAny) ToFloat64() float64 {
	return any.val
}

func (any *floatAny) ToString() string {
	return strconv.FormatFloat(any.val, 'E', -1, 64)
}

func (any *floatAny) WriteTo(stream *Stream) {
	stream.WriteFloat64(any.val)
}

func (any *floatAny) GetInterface() interface{} {
	return any.val
}

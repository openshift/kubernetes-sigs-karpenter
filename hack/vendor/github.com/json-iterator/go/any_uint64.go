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

type uint64Any struct {
	baseAny
	val uint64
}

func (any *uint64Any) LastError() error {
	return nil
}

func (any *uint64Any) ValueType() ValueType {
	return NumberValue
}

func (any *uint64Any) MustBeValid() Any {
	return any
}

func (any *uint64Any) ToBool() bool {
	return any.val != 0
}

func (any *uint64Any) ToInt() int {
	return int(any.val)
}

func (any *uint64Any) ToInt32() int32 {
	return int32(any.val)
}

func (any *uint64Any) ToInt64() int64 {
	return int64(any.val)
}

func (any *uint64Any) ToUint() uint {
	return uint(any.val)
}

func (any *uint64Any) ToUint32() uint32 {
	return uint32(any.val)
}

func (any *uint64Any) ToUint64() uint64 {
	return any.val
}

func (any *uint64Any) ToFloat32() float32 {
	return float32(any.val)
}

func (any *uint64Any) ToFloat64() float64 {
	return float64(any.val)
}

func (any *uint64Any) ToString() string {
	return strconv.FormatUint(any.val, 10)
}

func (any *uint64Any) WriteTo(stream *Stream) {
	stream.WriteUint64(any.val)
}

func (any *uint64Any) Parse() *Iterator {
	return nil
}

func (any *uint64Any) GetInterface() interface{} {
	return any.val
}

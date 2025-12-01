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

type nilAny struct {
	baseAny
}

func (any *nilAny) LastError() error {
	return nil
}

func (any *nilAny) ValueType() ValueType {
	return NilValue
}

func (any *nilAny) MustBeValid() Any {
	return any
}

func (any *nilAny) ToBool() bool {
	return false
}

func (any *nilAny) ToInt() int {
	return 0
}

func (any *nilAny) ToInt32() int32 {
	return 0
}

func (any *nilAny) ToInt64() int64 {
	return 0
}

func (any *nilAny) ToUint() uint {
	return 0
}

func (any *nilAny) ToUint32() uint32 {
	return 0
}

func (any *nilAny) ToUint64() uint64 {
	return 0
}

func (any *nilAny) ToFloat32() float32 {
	return 0
}

func (any *nilAny) ToFloat64() float64 {
	return 0
}

func (any *nilAny) ToString() string {
	return ""
}

func (any *nilAny) WriteTo(stream *Stream) {
	stream.WriteNil()
}

func (any *nilAny) Parse() *Iterator {
	return nil
}

func (any *nilAny) GetInterface() interface{} {
	return nil
}

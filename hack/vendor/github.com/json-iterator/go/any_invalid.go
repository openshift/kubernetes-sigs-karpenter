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

import "fmt"

type invalidAny struct {
	baseAny
	err error
}

func newInvalidAny(path []interface{}) *invalidAny {
	return &invalidAny{baseAny{}, fmt.Errorf("%v not found", path)}
}

func (any *invalidAny) LastError() error {
	return any.err
}

func (any *invalidAny) ValueType() ValueType {
	return InvalidValue
}

func (any *invalidAny) MustBeValid() Any {
	panic(any.err)
}

func (any *invalidAny) ToBool() bool {
	return false
}

func (any *invalidAny) ToInt() int {
	return 0
}

func (any *invalidAny) ToInt32() int32 {
	return 0
}

func (any *invalidAny) ToInt64() int64 {
	return 0
}

func (any *invalidAny) ToUint() uint {
	return 0
}

func (any *invalidAny) ToUint32() uint32 {
	return 0
}

func (any *invalidAny) ToUint64() uint64 {
	return 0
}

func (any *invalidAny) ToFloat32() float32 {
	return 0
}

func (any *invalidAny) ToFloat64() float64 {
	return 0
}

func (any *invalidAny) ToString() string {
	return ""
}

func (any *invalidAny) WriteTo(stream *Stream) {
}

func (any *invalidAny) Get(path ...interface{}) Any {
	if any.err == nil {
		return &invalidAny{baseAny{}, fmt.Errorf("get %v from invalid", path)}
	}
	return &invalidAny{baseAny{}, fmt.Errorf("%v, get %v from invalid", any.err, path)}
}

func (any *invalidAny) Parse() *Iterator {
	return nil
}

func (any *invalidAny) GetInterface() interface{} {
	return nil
}

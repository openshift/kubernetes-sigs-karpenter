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

package reflect2

import (
	"reflect"
	"unsafe"
)

type safeField struct {
	reflect.StructField
}

func (field *safeField) Offset() uintptr {
	return field.StructField.Offset
}

func (field *safeField) Name() string {
	return field.StructField.Name
}

func (field *safeField) PkgPath() string {
	return field.StructField.PkgPath
}

func (field *safeField) Type() Type {
	panic("not implemented")
}

func (field *safeField) Tag() reflect.StructTag {
	return field.StructField.Tag
}

func (field *safeField) Index() []int {
	return field.StructField.Index
}

func (field *safeField) Anonymous() bool {
	return field.StructField.Anonymous
}

func (field *safeField) Set(obj interface{}, value interface{}) {
	val := reflect.ValueOf(obj).Elem()
	val.FieldByIndex(field.Index()).Set(reflect.ValueOf(value).Elem())
}

func (field *safeField) UnsafeSet(obj unsafe.Pointer, value unsafe.Pointer) {
	panic("unsafe operation is not supported")
}

func (field *safeField) Get(obj interface{}) interface{} {
	val := reflect.ValueOf(obj).Elem().FieldByIndex(field.Index())
	ptr := reflect.New(val.Type())
	ptr.Elem().Set(val)
	return ptr.Interface()
}

func (field *safeField) UnsafeGet(obj unsafe.Pointer) unsafe.Pointer {
	panic("does not support unsafe operation")
}

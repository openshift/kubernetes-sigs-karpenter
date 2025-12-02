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

package decoder

import (
	"reflect"
	"unsafe"

	"github.com/goccy/go-json/internal/errors"
	"github.com/goccy/go-json/internal/runtime"
)

type invalidDecoder struct {
	typ        *runtime.Type
	kind       reflect.Kind
	structName string
	fieldName  string
}

func newInvalidDecoder(typ *runtime.Type, structName, fieldName string) *invalidDecoder {
	return &invalidDecoder{
		typ:        typ,
		kind:       typ.Kind(),
		structName: structName,
		fieldName:  fieldName,
	}
}

func (d *invalidDecoder) DecodeStream(s *Stream, depth int64, p unsafe.Pointer) error {
	return &errors.UnmarshalTypeError{
		Value:  "object",
		Type:   runtime.RType2Type(d.typ),
		Offset: s.totalOffset(),
		Struct: d.structName,
		Field:  d.fieldName,
	}
}

func (d *invalidDecoder) Decode(ctx *RuntimeContext, cursor, depth int64, p unsafe.Pointer) (int64, error) {
	return 0, &errors.UnmarshalTypeError{
		Value:  "object",
		Type:   runtime.RType2Type(d.typ),
		Offset: cursor,
		Struct: d.structName,
		Field:  d.fieldName,
	}
}

func (d *invalidDecoder) DecodePath(ctx *RuntimeContext, cursor, depth int64) ([][]byte, int64, error) {
	return nil, 0, &errors.UnmarshalTypeError{
		Value:  "object",
		Type:   runtime.RType2Type(d.typ),
		Offset: cursor,
		Struct: d.structName,
		Field:  d.fieldName,
	}
}

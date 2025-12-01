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
	"fmt"
	"unsafe"

	"github.com/goccy/go-json/internal/runtime"
)

type ptrDecoder struct {
	dec        Decoder
	typ        *runtime.Type
	structName string
	fieldName  string
}

func newPtrDecoder(dec Decoder, typ *runtime.Type, structName, fieldName string) *ptrDecoder {
	return &ptrDecoder{
		dec:        dec,
		typ:        typ,
		structName: structName,
		fieldName:  fieldName,
	}
}

func (d *ptrDecoder) contentDecoder() Decoder {
	dec, ok := d.dec.(*ptrDecoder)
	if !ok {
		return d.dec
	}
	return dec.contentDecoder()
}

//nolint:golint
//go:linkname unsafe_New reflect.unsafe_New
func unsafe_New(*runtime.Type) unsafe.Pointer

func UnsafeNew(t *runtime.Type) unsafe.Pointer {
	return unsafe_New(t)
}

func (d *ptrDecoder) DecodeStream(s *Stream, depth int64, p unsafe.Pointer) error {
	if s.skipWhiteSpace() == nul {
		s.read()
	}
	if s.char() == 'n' {
		if err := nullBytes(s); err != nil {
			return err
		}
		*(*unsafe.Pointer)(p) = nil
		return nil
	}
	var newptr unsafe.Pointer
	if *(*unsafe.Pointer)(p) == nil {
		newptr = unsafe_New(d.typ)
		*(*unsafe.Pointer)(p) = newptr
	} else {
		newptr = *(*unsafe.Pointer)(p)
	}
	if err := d.dec.DecodeStream(s, depth, newptr); err != nil {
		return err
	}
	return nil
}

func (d *ptrDecoder) Decode(ctx *RuntimeContext, cursor, depth int64, p unsafe.Pointer) (int64, error) {
	buf := ctx.Buf
	cursor = skipWhiteSpace(buf, cursor)
	if buf[cursor] == 'n' {
		if err := validateNull(buf, cursor); err != nil {
			return 0, err
		}
		if p != nil {
			*(*unsafe.Pointer)(p) = nil
		}
		cursor += 4
		return cursor, nil
	}
	var newptr unsafe.Pointer
	if *(*unsafe.Pointer)(p) == nil {
		newptr = unsafe_New(d.typ)
		*(*unsafe.Pointer)(p) = newptr
	} else {
		newptr = *(*unsafe.Pointer)(p)
	}
	c, err := d.dec.Decode(ctx, cursor, depth, newptr)
	if err != nil {
		*(*unsafe.Pointer)(p) = nil
		return 0, err
	}
	cursor = c
	return cursor, nil
}

func (d *ptrDecoder) DecodePath(ctx *RuntimeContext, cursor, depth int64) ([][]byte, int64, error) {
	return nil, 0, fmt.Errorf("json: ptr decoder does not support decode path")
}

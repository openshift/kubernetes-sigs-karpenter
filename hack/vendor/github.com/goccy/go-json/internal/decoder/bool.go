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

	"github.com/goccy/go-json/internal/errors"
)

type boolDecoder struct {
	structName string
	fieldName  string
}

func newBoolDecoder(structName, fieldName string) *boolDecoder {
	return &boolDecoder{structName: structName, fieldName: fieldName}
}

func (d *boolDecoder) DecodeStream(s *Stream, depth int64, p unsafe.Pointer) error {
	c := s.skipWhiteSpace()
	for {
		switch c {
		case 't':
			if err := trueBytes(s); err != nil {
				return err
			}
			**(**bool)(unsafe.Pointer(&p)) = true
			return nil
		case 'f':
			if err := falseBytes(s); err != nil {
				return err
			}
			**(**bool)(unsafe.Pointer(&p)) = false
			return nil
		case 'n':
			if err := nullBytes(s); err != nil {
				return err
			}
			return nil
		case nul:
			if s.read() {
				c = s.char()
				continue
			}
			goto ERROR
		}
		break
	}
ERROR:
	return errors.ErrUnexpectedEndOfJSON("bool", s.totalOffset())
}

func (d *boolDecoder) Decode(ctx *RuntimeContext, cursor, depth int64, p unsafe.Pointer) (int64, error) {
	buf := ctx.Buf
	cursor = skipWhiteSpace(buf, cursor)
	switch buf[cursor] {
	case 't':
		if err := validateTrue(buf, cursor); err != nil {
			return 0, err
		}
		cursor += 4
		**(**bool)(unsafe.Pointer(&p)) = true
		return cursor, nil
	case 'f':
		if err := validateFalse(buf, cursor); err != nil {
			return 0, err
		}
		cursor += 5
		**(**bool)(unsafe.Pointer(&p)) = false
		return cursor, nil
	case 'n':
		if err := validateNull(buf, cursor); err != nil {
			return 0, err
		}
		cursor += 4
		return cursor, nil
	}
	return 0, errors.ErrUnexpectedEndOfJSON("bool", cursor)
}

func (d *boolDecoder) DecodePath(ctx *RuntimeContext, cursor, depth int64) ([][]byte, int64, error) {
	return nil, 0, fmt.Errorf("json: bool decoder does not support decode path")
}

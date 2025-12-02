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
	"encoding/json"
	"github.com/modern-go/reflect2"
	"unsafe"
)

var jsonRawMessageType = reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()
var jsoniterRawMessageType = reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()

func createEncoderOfJsonRawMessage(ctx *ctx, typ reflect2.Type) ValEncoder {
	if typ == jsonRawMessageType {
		return &jsonRawMessageCodec{}
	}
	if typ == jsoniterRawMessageType {
		return &jsoniterRawMessageCodec{}
	}
	return nil
}

func createDecoderOfJsonRawMessage(ctx *ctx, typ reflect2.Type) ValDecoder {
	if typ == jsonRawMessageType {
		return &jsonRawMessageCodec{}
	}
	if typ == jsoniterRawMessageType {
		return &jsoniterRawMessageCodec{}
	}
	return nil
}

type jsonRawMessageCodec struct {
}

func (codec *jsonRawMessageCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if iter.ReadNil() {
		*((*json.RawMessage)(ptr)) = nil
	} else {
		*((*json.RawMessage)(ptr)) = iter.SkipAndReturnBytes()
	}
}

func (codec *jsonRawMessageCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	if *((*json.RawMessage)(ptr)) == nil {
		stream.WriteNil()
	} else {
		stream.WriteRaw(string(*((*json.RawMessage)(ptr))))
	}
}

func (codec *jsonRawMessageCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(*((*json.RawMessage)(ptr))) == 0
}

type jsoniterRawMessageCodec struct {
}

func (codec *jsoniterRawMessageCodec) Decode(ptr unsafe.Pointer, iter *Iterator) {
	if iter.ReadNil() {
		*((*RawMessage)(ptr)) = nil
	} else {
		*((*RawMessage)(ptr)) = iter.SkipAndReturnBytes()
	}
}

func (codec *jsoniterRawMessageCodec) Encode(ptr unsafe.Pointer, stream *Stream) {
	if *((*RawMessage)(ptr)) == nil {
		stream.WriteNil()
	} else {
		stream.WriteRaw(string(*((*RawMessage)(ptr))))
	}
}

func (codec *jsoniterRawMessageCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(*((*RawMessage)(ptr))) == 0
}

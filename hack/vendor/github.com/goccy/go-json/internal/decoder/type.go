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
	"context"
	"encoding"
	"encoding/json"
	"reflect"
	"unsafe"
)

type Decoder interface {
	Decode(*RuntimeContext, int64, int64, unsafe.Pointer) (int64, error)
	DecodePath(*RuntimeContext, int64, int64) ([][]byte, int64, error)
	DecodeStream(*Stream, int64, unsafe.Pointer) error
}

const (
	nul                   = '\000'
	maxDecodeNestingDepth = 10000
)

type unmarshalerContext interface {
	UnmarshalJSON(context.Context, []byte) error
}

var (
	unmarshalJSONType        = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	unmarshalJSONContextType = reflect.TypeOf((*unmarshalerContext)(nil)).Elem()
	unmarshalTextType        = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
)

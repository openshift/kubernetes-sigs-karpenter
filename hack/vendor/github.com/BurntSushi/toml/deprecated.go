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

package toml

import (
	"encoding"
	"io"
)

// TextMarshaler is an alias for encoding.TextMarshaler.
//
// Deprecated: use encoding.TextMarshaler
type TextMarshaler encoding.TextMarshaler

// TextUnmarshaler is an alias for encoding.TextUnmarshaler.
//
// Deprecated: use encoding.TextUnmarshaler
type TextUnmarshaler encoding.TextUnmarshaler

// DecodeReader is an alias for NewDecoder(r).Decode(v).
//
// Deprecated: use NewDecoder(reader).Decode(&value).
func DecodeReader(r io.Reader, v any) (MetaData, error) { return NewDecoder(r).Decode(v) }

// PrimitiveDecode is an alias for MetaData.PrimitiveDecode().
//
// Deprecated: use MetaData.PrimitiveDecode.
func PrimitiveDecode(primValue Primitive, v any) error {
	md := MetaData{decoded: make(map[string]struct{})}
	return md.unify(primValue.undecoded, rvalue(v))
}

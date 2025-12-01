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

package jsonschema

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
)

// Decoder specifies how to decode specific contentEncoding.
type Decoder struct {
	// Name of contentEncoding.
	Name string
	// Decode given string to byte array.
	Decode func(string) ([]byte, error)
}

var decoders = map[string]*Decoder{
	"base64": {
		Name: "base64",
		Decode: func(s string) ([]byte, error) {
			return base64.StdEncoding.DecodeString(s)
		},
	},
}

// MediaType specified how to validate bytes against specific contentMediaType.
type MediaType struct {
	// Name of contentMediaType.
	Name string

	// Validate checks whether bytes conform to this mediatype.
	Validate func([]byte) error

	// UnmarshalJSON unmarshals bytes into json value.
	// This must be nil if this mediatype is not compatible
	// with json.
	UnmarshalJSON func([]byte) (any, error)
}

var mediaTypes = map[string]*MediaType{
	"application/json": {
		Name: "application/json",
		Validate: func(b []byte) error {
			var v any
			return json.Unmarshal(b, &v)
		},
		UnmarshalJSON: func(b []byte) (any, error) {
			return UnmarshalJSON(bytes.NewReader(b))
		},
	},
}

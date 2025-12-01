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

// Package document provides interface definitions and error types for document types.
//
// A document is a protocol-agnostic type which supports a JSON-like data-model. You can use this type to send
// UTF-8 strings, arbitrary precision numbers, booleans, nulls, a list of these values, and a map of UTF-8
// strings to these values.
//
// API Clients expose document constructors in their respective client document packages which must be used to
// Marshal and Unmarshal Go types to and from their respective protocol representations.
//
// See the Marshaler and Unmarshaler type documentation for more details on how to Go types can be converted to and from
// document types.
package document

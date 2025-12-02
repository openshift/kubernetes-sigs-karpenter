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

// Package toml implements decoding and encoding of TOML files.
//
// This package supports TOML v1.0.0, as specified at https://toml.io
//
// The github.com/BurntSushi/toml/cmd/tomlv package implements a TOML validator,
// and can be used to verify if TOML document is valid. It can also be used to
// print the type of each key.
package toml

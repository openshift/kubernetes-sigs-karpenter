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

// Package toml is a TOML parser and manipulation library.
//
// This version supports the specification as described in
// https://github.com/toml-lang/toml/blob/master/versions/en/toml-v0.5.0.md
//
// Marshaling
//
// Go-toml can marshal and unmarshal TOML documents from and to data
// structures.
//
// TOML document as a tree
//
// Go-toml can operate on a TOML document as a tree. Use one of the Load*
// functions to parse TOML data and obtain a Tree instance, then one of its
// methods to manipulate the tree.
//
// JSONPath-like queries
//
// The package github.com/pelletier/go-toml/query implements a system
// similar to JSONPath to quickly retrieve elements of a TOML document using a
// single expression. See the package documentation for more information.
//
package toml

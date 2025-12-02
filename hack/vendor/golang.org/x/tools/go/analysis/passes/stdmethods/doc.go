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

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package stdmethods defines an Analyzer that checks for misspellings
// in the signatures of methods similar to well-known interfaces.
//
// # Analyzer stdmethods
//
// stdmethods: check signature of methods of well-known interfaces
//
// Sometimes a type may be intended to satisfy an interface but may fail to
// do so because of a mistake in its method signature.
// For example, the result of this WriteTo method should be (int64, error),
// not error, to satisfy io.WriterTo:
//
//	type myWriterTo struct{...}
//	func (myWriterTo) WriteTo(w io.Writer) error { ... }
//
// This check ensures that each method whose name matches one of several
// well-known interface methods from the standard library has the correct
// signature for that interface.
//
// Checked method names include:
//
//	Format GobEncode GobDecode MarshalJSON MarshalXML
//	Peek ReadByte ReadFrom ReadRune Scan Seek
//	UnmarshalJSON UnreadByte UnreadRune WriteByte
//	WriteTo
package stdmethods

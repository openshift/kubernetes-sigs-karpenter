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

// Package slog defines an Analyzer that checks for
// mismatched key-value pairs in log/slog calls.
//
// # Analyzer slog
//
// slog: check for invalid structured logging calls
//
// The slog checker looks for calls to functions from the log/slog
// package that take alternating key-value pairs. It reports calls
// where an argument in a key position is neither a string nor a
// slog.Attr, and where a final key is missing its value.
// For example,it would report
//
//	slog.Warn("message", 11, "k") // slog.Warn arg "11" should be a string or a slog.Attr
//
// and
//
//	slog.Info("message", "k1", v1, "k2") // call to slog.Info missing a final value
package slog

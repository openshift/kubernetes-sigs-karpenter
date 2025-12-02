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

// Package lostcancel defines an Analyzer that checks for failure to
// call a context cancellation function.
//
// # Analyzer lostcancel
//
// lostcancel: check cancel func returned by context.WithCancel is called
//
// The cancellation function returned by context.WithCancel, WithTimeout,
// WithDeadline and variants such as WithCancelCause must be called,
// or the new context will remain live until its parent context is cancelled.
// (The background context is never cancelled.)
package lostcancel

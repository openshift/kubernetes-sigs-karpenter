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

// Package sigchanyzer defines an Analyzer that detects
// misuse of unbuffered signal as argument to signal.Notify.
//
// # Analyzer sigchanyzer
//
// sigchanyzer: check for unbuffered channel of os.Signal
//
// This checker reports call expression of the form
//
//	signal.Notify(c <-chan os.Signal, sig ...os.Signal),
//
// where c is an unbuffered channel, which can be at risk of missing the signal.
package sigchanyzer

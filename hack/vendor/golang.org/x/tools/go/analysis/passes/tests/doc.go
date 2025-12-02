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

// Package tests defines an Analyzer that checks for common mistaken
// usages of tests and examples.
//
// # Analyzer tests
//
// tests: check for common mistaken usages of tests and examples
//
// The tests checker walks Test, Benchmark, Fuzzing and Example functions checking
// malformed names, wrong signatures and examples documenting non-existent
// identifiers.
//
// Please see the documentation for package testing in golang.org/pkg/testing
// for the conventions that are enforced for Tests, Benchmarks, and Examples.
package tests

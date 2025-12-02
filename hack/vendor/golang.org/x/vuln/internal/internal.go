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

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package internal contains functionality for x/vuln.
package internal

// IDDirectory is the name of the directory that contains entries
// listed by their IDs.
const IDDirectory = "ID"

// Pseudo-module paths used for parts of the Go system.
// These are technically not valid module paths, so we
// mustn't pass them to module.EscapePath.
// Keep in sync with vulndb/internal/database/generate.go.
const (
	// GoStdModulePath is the internal Go module path string used
	// when listing vulnerabilities in standard library.
	GoStdModulePath = "stdlib"

	// GoCmdModulePath is the internal Go module path string used
	// when listing vulnerabilities in the go command.
	GoCmdModulePath = "toolchain"

	// UnknownModulePath is a special module path for when we cannot work out
	// the module for a package.
	UnknownModulePath = "unknown-module"

	// UnknownPackagePath is a special package path for when we cannot work out
	// the packagUnknownModulePath = "unknown"
	UnknownPackagePath = "unknown-package"
)

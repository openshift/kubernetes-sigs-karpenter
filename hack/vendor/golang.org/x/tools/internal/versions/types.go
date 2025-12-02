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

package versions

import (
	"go/ast"
	"go/types"
)

// FileVersion returns a file's Go version.
// The reported version is an unknown Future version if a
// version cannot be determined.
func FileVersion(info *types.Info, file *ast.File) string {
	// In tools built with Go >= 1.22, the Go version of a file
	// follow a cascades of sources:
	// 1) types.Info.FileVersion, which follows the cascade:
	//   1.a) file version (ast.File.GoVersion),
	//   1.b) the package version (types.Config.GoVersion), or
	// 2) is some unknown Future version.
	//
	// File versions require a valid package version to be provided to types
	// in Config.GoVersion. Config.GoVersion is either from the package's module
	// or the toolchain (go run). This value should be provided by go/packages
	// or unitchecker.Config.GoVersion.
	if v := info.FileVersions[file]; IsValid(v) {
		return v
	}
	// Note: we could instead return runtime.Version() [if valid].
	// This would act as a max version on what a tool can support.
	return Future
}

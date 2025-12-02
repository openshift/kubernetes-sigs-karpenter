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

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssautil

// This file contains deprecated public APIs.
// We discourage their use.

import (
	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
)

// CreateProgram returns a new program in SSA form, given a program
// loaded from source.  An SSA package is created for each transitively
// error-free package of lprog.
//
// Code for bodies of functions is not built until Build is called
// on the result.
//
// The mode parameter controls diagnostics and checking during SSA construction.
//
// Deprecated: Use [golang.org/x/tools/go/packages] and the [Packages]
// function instead; see ssa.Example_loadPackages.
func CreateProgram(lprog *loader.Program, mode ssa.BuilderMode) *ssa.Program {
	prog := ssa.NewProgram(lprog.Fset, mode)

	for _, info := range lprog.AllPackages {
		if info.TransitivelyErrorFree {
			prog.CreatePackage(info.Pkg, info.Files, &info.Info, info.Importable)
		}
	}

	return prog
}

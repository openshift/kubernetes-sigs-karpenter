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

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeutil

import "go/types"

// Dependencies returns all dependencies of the specified packages.
//
// Dependent packages appear in topological order: if package P imports
// package Q, Q appears earlier than P in the result.
// The algorithm follows import statements in the order they
// appear in the source code, so the result is a total order.
func Dependencies(pkgs ...*types.Package) []*types.Package {
	var result []*types.Package
	seen := make(map[*types.Package]bool)
	var visit func(pkgs []*types.Package)
	visit = func(pkgs []*types.Package) {
		for _, p := range pkgs {
			if !seen[p] {
				seen[p] = true
				visit(p.Imports())
				result = append(result, p)
			}
		}
	}
	visit(pkgs)
	return result
}

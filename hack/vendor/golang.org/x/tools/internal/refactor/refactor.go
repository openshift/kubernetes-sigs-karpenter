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

// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package refactor provides operators to compute common textual edits
// for refactoring tools.
//
// This package should not use features of the analysis API
// other than [analysis.TextEdit].
package refactor

import (
	"fmt"
	"go/token"
	"go/types"
)

// FreshName returns the name of an identifier that is undefined
// at the specified position, based on the preferred name.
func FreshName(scope *types.Scope, pos token.Pos, preferred string) string {
	newName := preferred
	for i := 0; ; i++ {
		if _, obj := scope.LookupParent(newName, pos); obj == nil {
			break // fresh
		}
		newName = fmt.Sprintf("%s%d", preferred, i)
	}
	return newName
}

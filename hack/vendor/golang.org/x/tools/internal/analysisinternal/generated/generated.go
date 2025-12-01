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

// Package generated defines an analyzer whose result makes it
// convenient to skip diagnostics within generated files.
package generated

import (
	"go/ast"
	"go/token"
	"reflect"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:       "generated",
	Doc:        "detect which Go files are generated",
	URL:        "https://pkg.go.dev/golang.org/x/tools/internal/analysisinternal/generated",
	ResultType: reflect.TypeFor[*Result](),
	Run: func(pass *analysis.Pass) (any, error) {
		set := make(map[*token.File]bool)
		for _, file := range pass.Files {
			if ast.IsGenerated(file) {
				set[pass.Fset.File(file.FileStart)] = true
			}
		}
		return &Result{fset: pass.Fset, generatedFiles: set}, nil
	},
}

type Result struct {
	fset           *token.FileSet
	generatedFiles map[*token.File]bool
}

// IsGenerated reports whether the position is within a generated file.
func (r *Result) IsGenerated(pos token.Pos) bool {
	return r.generatedFiles[r.fset.File(pos)]
}

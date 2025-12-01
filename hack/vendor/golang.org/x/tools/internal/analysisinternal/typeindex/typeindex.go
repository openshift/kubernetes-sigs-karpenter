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

// Package typeindex defines an analyzer that provides a
// [golang.org/x/tools/internal/typesinternal/typeindex.Index].
//
// Like [golang.org/x/tools/go/analysis/passes/inspect], it is
// intended to be used as a helper by other analyzers; it reports no
// diagnostics of its own.
package typeindex

import (
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/internal/typesinternal/typeindex"
)

var Analyzer = &analysis.Analyzer{
	Name: "typeindex",
	Doc:  "indexes of type information for later passes",
	URL:  "https://pkg.go.dev/golang.org/x/tools/internal/analysisinternal/typeindex",
	Run: func(pass *analysis.Pass) (any, error) {
		inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
		return typeindex.New(inspect, pass.Pkg, pass.TypesInfo), nil
	},
	RunDespiteErrors: true,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
	ResultType:       reflect.TypeOf(new(typeindex.Index)),
}

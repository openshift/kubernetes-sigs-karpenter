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

// Package err113 is a Golang linter to check the errors handling expressions
package err113

import (
	"go/ast"
	"go/printer"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// NewAnalyzer creates a new analysis.Analyzer instance tuned to run err113 checks.
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "err113",
		Doc:  "checks the error handling rules according to the Go 1.13 new error type",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		tlds := enumerateFileDecls(file)

		ast.Inspect(
			file,
			func(n ast.Node) bool {
				return inspectComparision(file, pass, n) &&
					inspectDefinition(pass, tlds, n)
			},
		)
	}

	return nil, nil
}

// render returns the pretty-print of the given node.
func render(fset *token.FileSet, x any) string {
	var buf strings.Builder
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}

	return buf.String()
}

func enumerateFileDecls(f *ast.File) map[*ast.CallExpr]struct{} {
	res := make(map[*ast.CallExpr]struct{})

	var ces []*ast.CallExpr // nolint: prealloc

	for _, d := range f.Decls {
		ces = append(ces, enumerateDeclVars(d)...)
	}

	for _, ce := range ces {
		res[ce] = struct{}{}
	}

	return res
}

func enumerateDeclVars(d ast.Decl) (res []*ast.CallExpr) {
	td, ok := d.(*ast.GenDecl)
	if !ok || td.Tok != token.VAR {
		return nil
	}

	for _, s := range td.Specs {
		res = append(res, enumerateSpecValues(s)...)
	}

	return res
}

func enumerateSpecValues(s ast.Spec) (res []*ast.CallExpr) {
	vs, ok := s.(*ast.ValueSpec)
	if !ok {
		return nil
	}

	for _, v := range vs.Values {
		if ce, ok := v.(*ast.CallExpr); ok {
			res = append(res, ce)
		}
	}

	return res
}

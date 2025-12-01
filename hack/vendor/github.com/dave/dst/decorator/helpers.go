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

package decorator

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"

	"github.com/dave/dst"
)

// Parse uses parser.ParseFile to parse and decorate a Go source file. The src parameter should
// be string, []byte, or io.Reader.
func Parse(src interface{}) (*dst.File, error) {
	return NewDecorator(token.NewFileSet()).Parse(src)
}

// ParseFile uses parser.ParseFile to parse and decorate a Go source file. The ParseComments flag is
// added to mode if it doesn't exist.
func ParseFile(fset *token.FileSet, filename string, src interface{}, mode parser.Mode) (*dst.File, error) {
	return NewDecorator(fset).ParseFile(filename, src, mode)
}

// ParseDir uses parser.ParseDir to parse and decorate a directory containing Go source. The
// ParseComments flag is added to mode if it doesn't exist.
func ParseDir(fset *token.FileSet, dir string, filter func(os.FileInfo) bool, mode parser.Mode) (map[string]*dst.Package, error) {
	return NewDecorator(fset).ParseDir(dir, filter, mode)
}

// Decorate decorates an ast.Node and returns a dst.Node.
func Decorate(fset *token.FileSet, n ast.Node) (dst.Node, error) {
	return NewDecorator(fset).DecorateNode(n)
}

// Decorate decorates a *ast.File and returns a *dst.File.
func DecorateFile(fset *token.FileSet, f *ast.File) (*dst.File, error) {
	return NewDecorator(fset).DecorateFile(f)
}

// Print uses format.Node to print a *dst.File to stdout
func Print(f *dst.File) error {
	return Fprint(os.Stdout, f)
}

// Fprint uses format.Node to print a *dst.File to a writer
func Fprint(w io.Writer, f *dst.File) error {
	fset, af, err := RestoreFile(f)
	if err != nil {
		return err
	}
	return format.Node(w, fset, af)
}

// RestoreFile restores a *dst.File to a *token.FileSet and a *ast.File
func RestoreFile(file *dst.File) (*token.FileSet, *ast.File, error) {
	r := NewRestorer()
	f, err := r.RestoreFile(file)
	if err != nil {
		return nil, nil, err
	}
	return r.Fset, f, nil
}

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

package formatter

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

type GoFmtFormatter struct {
	fset *token.FileSet
}

func NewGoFmtFormatter(fset *token.FileSet) *GoFmtFormatter {
	return &GoFmtFormatter{fset: fset}
}

func (f GoFmtFormatter) Format(exp ast.Expr) string {
	var buf bytes.Buffer
	_ = printer.Fprint(&buf, f.fset, exp)
	return buf.String()
}

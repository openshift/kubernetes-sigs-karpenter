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

package analysisutil

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
)

// NodeString is a more powerful analogue of types.ExprString.
// Return empty string if node AST is invalid.
func NodeString(fset *token.FileSet, node ast.Node) string {
	if v := formatNode(fset, node); v != nil {
		return v.String()
	}
	return ""
}

// NodeBytes works as NodeString but returns a byte slice.
// Return nil if node AST is invalid.
func NodeBytes(fset *token.FileSet, node ast.Node) []byte {
	if v := formatNode(fset, node); v != nil {
		return v.Bytes()
	}
	return nil
}

func formatNode(fset *token.FileSet, node ast.Node) *bytes.Buffer {
	buf := new(bytes.Buffer)
	if err := format.Node(buf, fset, node); err != nil {
		return nil
	}
	return buf
}

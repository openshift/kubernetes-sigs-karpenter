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

// Package strparse provides convenience wrappers around `go/parser` for simple
// expression, statement and declaration parsing from string.
//
// Can be used to construct AST nodes using source syntax.
package strparse

import (
	"go/ast"
	"go/parser"
	"go/token"
)

var (
	// BadExpr is returned as a parse result for malformed expressions.
	// Should be treated as constant or readonly variable.
	BadExpr = &ast.BadExpr{}

	// BadStmt is returned as a parse result for malformed statmenents.
	// Should be treated as constant or readonly variable.
	BadStmt = &ast.BadStmt{}

	// BadDecl is returned as a parse result for malformed declarations.
	// Should be treated as constant or readonly variable.
	BadDecl = &ast.BadDecl{}
)

// Expr parses single expression node from s.
// In case of parse error, BadExpr is returned.
func Expr(s string) ast.Expr {
	node, err := parser.ParseExpr(s)
	if err != nil {
		return BadExpr
	}
	return node
}

// Stmt parses single statement node from s.
// In case of parse error, BadStmt is returned.
func Stmt(s string) ast.Stmt {
	node, err := parser.ParseFile(token.NewFileSet(), "", "package main;func main() {"+s+"}", 0)
	if err != nil {
		return BadStmt
	}
	fn := node.Decls[0].(*ast.FuncDecl)
	if len(fn.Body.List) != 1 {
		return BadStmt
	}
	return fn.Body.List[0]
}

// Decl parses single declaration node from s.
// In case of parse error, BadDecl is returned.
func Decl(s string) ast.Decl {
	node, err := parser.ParseFile(token.NewFileSet(), "", "package main;"+s, 0)
	if err != nil || len(node.Decls) != 1 {
		return BadDecl
	}
	return node.Decls[0]
}

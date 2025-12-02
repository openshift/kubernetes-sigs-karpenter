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

package ast

type Stmt interface {
	PositionHolder
	stmtMarker()
}

type StmtBase struct {
	Node
}

func (stmt *StmtBase) stmtMarker() {}

type AssignStmt struct {
	StmtBase

	Lhs []Expr
	Rhs []Expr
}

type LocalAssignStmt struct {
	StmtBase

	Names []string
	Exprs []Expr
}

type FuncCallStmt struct {
	StmtBase

	Expr Expr
}

type DoBlockStmt struct {
	StmtBase

	Stmts []Stmt
}

type WhileStmt struct {
	StmtBase

	Condition Expr
	Stmts     []Stmt
}

type RepeatStmt struct {
	StmtBase

	Condition Expr
	Stmts     []Stmt
}

type IfStmt struct {
	StmtBase

	Condition Expr
	Then      []Stmt
	Else      []Stmt
}

type NumberForStmt struct {
	StmtBase

	Name  string
	Init  Expr
	Limit Expr
	Step  Expr
	Stmts []Stmt
}

type GenericForStmt struct {
	StmtBase

	Names []string
	Exprs []Expr
	Stmts []Stmt
}

type FuncDefStmt struct {
	StmtBase

	Name *FuncName
	Func *FunctionExpr
}

type ReturnStmt struct {
	StmtBase

	Exprs []Expr
}

type BreakStmt struct {
	StmtBase
}

type LabelStmt struct {
	StmtBase

	Name string
}

type GotoStmt struct {
	StmtBase

	Label string
}

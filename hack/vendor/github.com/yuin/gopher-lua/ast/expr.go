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

type Expr interface {
	PositionHolder
	exprMarker()
}

type ExprBase struct {
	Node
}

func (expr *ExprBase) exprMarker() {}

/* ConstExprs {{{ */

type ConstExpr interface {
	Expr
	constExprMarker()
}

type ConstExprBase struct {
	ExprBase
}

func (expr *ConstExprBase) constExprMarker() {}

type TrueExpr struct {
	ConstExprBase
}

type FalseExpr struct {
	ConstExprBase
}

type NilExpr struct {
	ConstExprBase
}

type NumberExpr struct {
	ConstExprBase

	Value string
}

type StringExpr struct {
	ConstExprBase

	Value string
}

/* ConstExprs }}} */

type Comma3Expr struct {
	ExprBase
	AdjustRet bool
}

type IdentExpr struct {
	ExprBase

	Value string
}

type AttrGetExpr struct {
	ExprBase

	Object Expr
	Key    Expr
}

type TableExpr struct {
	ExprBase

	Fields []*Field
}

type FuncCallExpr struct {
	ExprBase

	Func      Expr
	Receiver  Expr
	Method    string
	Args      []Expr
	AdjustRet bool
}

type LogicalOpExpr struct {
	ExprBase

	Operator string
	Lhs      Expr
	Rhs      Expr
}

type RelationalOpExpr struct {
	ExprBase

	Operator string
	Lhs      Expr
	Rhs      Expr
}

type StringConcatOpExpr struct {
	ExprBase

	Lhs Expr
	Rhs Expr
}

type ArithmeticOpExpr struct {
	ExprBase

	Operator string
	Lhs      Expr
	Rhs      Expr
}

type UnaryMinusOpExpr struct {
	ExprBase
	Expr Expr
}

type UnaryNotOpExpr struct {
	ExprBase
	Expr Expr
}

type UnaryLenOpExpr struct {
	ExprBase
	Expr Expr
}

type FunctionExpr struct {
	ExprBase

	ParList *ParList
	Stmts   []Stmt
}

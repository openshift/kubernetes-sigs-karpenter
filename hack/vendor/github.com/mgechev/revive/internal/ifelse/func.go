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

package ifelse

import (
	"fmt"
	"go/ast"
)

// Call contains the name of a function that deviates control flow.
type Call struct {
	Pkg  string // The package qualifier of the function, if not built-in.
	Name string // The function name.
}

// DeviatingFuncs lists known control flow deviating function calls.
var DeviatingFuncs = map[Call]BranchKind{
	{"os", "Exit"}:     Exit,
	{"log", "Fatal"}:   Exit,
	{"log", "Fatalf"}:  Exit,
	{"log", "Fatalln"}: Exit,
	{"", "panic"}:      Panic,
	{"log", "Panic"}:   Panic,
	{"log", "Panicf"}:  Panic,
	{"log", "Panicln"}: Panic,
}

// ExprCall gets the Call of an ExprStmt, if any.
func ExprCall(expr *ast.ExprStmt) (Call, bool) {
	call, ok := expr.X.(*ast.CallExpr)
	if !ok {
		return Call{}, false
	}
	switch v := call.Fun.(type) {
	case *ast.Ident:
		return Call{Name: v.Name}, true
	case *ast.SelectorExpr:
		if ident, ok := v.X.(*ast.Ident); ok {
			return Call{Name: v.Sel.Name, Pkg: ident.Name}, true
		}
	}
	return Call{}, false
}

// String returns the function name with package qualifier (if any).
func (f Call) String() string {
	if f.Pkg != "" {
		return fmt.Sprintf("%s.%s", f.Pkg, f.Name)
	}
	return f.Name
}

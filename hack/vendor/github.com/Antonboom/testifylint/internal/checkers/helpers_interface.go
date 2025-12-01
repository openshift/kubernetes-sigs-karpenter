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

package checkers

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"

	"github.com/Antonboom/testifylint/internal/analysisutil"
	"github.com/Antonboom/testifylint/internal/testify"
)

func isEmptyInterface(pass *analysis.Pass, expr ast.Expr) bool {
	t, ok := pass.TypesInfo.Types[expr]
	if !ok {
		return false
	}
	return isEmptyInterfaceType(t.Type)
}

func isEmptyInterfaceType(t types.Type) bool {
	iface, ok := t.Underlying().(*types.Interface)
	return ok && iface.NumMethods() == 0
}

func implementsTestifySuite(pass *analysis.Pass, e ast.Expr) bool {
	suiteIfaceObj := analysisutil.ObjectOf(pass.Pkg, testify.SuitePkgPath, "TestingSuite")
	return (suiteIfaceObj != nil) && implements(pass, e, suiteIfaceObj)
}

func implementsTestingT(pass *analysis.Pass, e ast.Expr) bool {
	return implementsAssertTestingT(pass, e) || implementsRequireTestingT(pass, e)
}

func implementsAssertTestingT(pass *analysis.Pass, e ast.Expr) bool {
	assertTestingTObj := analysisutil.ObjectOf(pass.Pkg, testify.AssertPkgPath, "TestingT")
	return (assertTestingTObj != nil) && implements(pass, e, assertTestingTObj)
}

func implementsRequireTestingT(pass *analysis.Pass, e ast.Expr) bool {
	requireTestingTObj := analysisutil.ObjectOf(pass.Pkg, testify.RequirePkgPath, "TestingT")
	return (requireTestingTObj != nil) && implements(pass, e, requireTestingTObj)
}

func implements(pass *analysis.Pass, e ast.Expr, ifaceObj types.Object) bool {
	t := pass.TypesInfo.TypeOf(e)
	if t == nil {
		return false
	}
	return types.Implements(t, ifaceObj.Type().Underlying().(*types.Interface))
}

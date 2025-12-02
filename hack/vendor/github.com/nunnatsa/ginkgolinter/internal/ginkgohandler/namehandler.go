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

package ginkgohandler

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	config "github.com/nunnatsa/ginkgolinter/config"
)

// nameHandler is used when importing ginkgo without name; i.e.
// import "github.com/onsi/ginkgo"
//
// or with a custom name; e.g.
// import customname "github.com/onsi/ginkgo"
type nameHandler string

func (h nameHandler) HandleGinkgoSpecs(expr ast.Expr, config config.Config, pass *analysis.Pass) bool {
	return handleGinkgoSpecs(expr, config, pass, h)
}

func (h nameHandler) getFocusContainerName(exp *ast.CallExpr) (bool, *ast.Ident) {
	if sel, ok := exp.Fun.(*ast.SelectorExpr); ok {
		if id, ok := sel.X.(*ast.Ident); ok && id.Name == string(h) {
			return isFocusContainer(sel.Sel.Name), sel.Sel
		}
	}
	return false, nil
}

func (h nameHandler) isWrapContainer(exp *ast.CallExpr) bool {
	if sel, ok := exp.Fun.(*ast.SelectorExpr); ok {
		if id, ok := sel.X.(*ast.Ident); ok && id.Name == string(h) {
			return isWrapContainer(sel.Sel.Name)
		}
	}

	return false
}

func (h nameHandler) isFocusSpec(exp ast.Expr) bool {
	if selExp, ok := exp.(*ast.SelectorExpr); ok {
		if x, ok := selExp.X.(*ast.Ident); ok && x.Name == string(h) {
			return selExp.Sel.Name == focusSpec
		}
	}

	return false
}

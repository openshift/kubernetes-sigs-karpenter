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

	"github.com/nunnatsa/ginkgolinter/config"
)

// dotHandler is used when importing ginkgo with dot; i.e.
// import . "github.com/onsi/ginkgo"
type dotHandler struct{}

func (h dotHandler) HandleGinkgoSpecs(expr ast.Expr, config config.Config, pass *analysis.Pass) bool {
	return handleGinkgoSpecs(expr, config, pass, h)
}

func (h dotHandler) getFocusContainerName(exp *ast.CallExpr) (bool, *ast.Ident) {
	if fun, ok := exp.Fun.(*ast.Ident); ok {
		return isFocusContainer(fun.Name), fun
	}
	return false, nil
}

func (h dotHandler) isWrapContainer(exp *ast.CallExpr) bool {
	if fun, ok := exp.Fun.(*ast.Ident); ok {
		return isWrapContainer(fun.Name)
	}
	return false
}

func (h dotHandler) isFocusSpec(exp ast.Expr) bool {
	id, ok := exp.(*ast.Ident)
	return ok && id.Name == focusSpec
}

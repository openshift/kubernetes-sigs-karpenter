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

package matcher

import (
	"go/ast"
	gotypes "go/types"

	"golang.org/x/tools/go/analysis"

	"github.com/nunnatsa/ginkgolinter/internal/gomegahandler"
)

type HaveValueMatcher struct {
	nested *Matcher
}

func (m *HaveValueMatcher) Type() Type {
	return HaveValueMatherType
}
func (m *HaveValueMatcher) MatcherName() string {
	return haveValue
}

func (m *HaveValueMatcher) GetNested() *Matcher {
	return m.nested
}

type WithTransformMatcher struct {
	funcType gotypes.Type
	nested   *Matcher
}

func (m *WithTransformMatcher) Type() Type {
	return WithTransformMatherType
}
func (m *WithTransformMatcher) MatcherName() string {
	return withTransform
}

func (m *WithTransformMatcher) GetNested() *Matcher {
	return m.nested
}

func (m *WithTransformMatcher) GetFuncType() gotypes.Type {
	return m.funcType
}

func getNestedMatcher(orig, clone *ast.CallExpr, offset int, pass *analysis.Pass, handler gomegahandler.Handler) (*Matcher, bool) {
	if origNested, ok := orig.Args[offset].(*ast.CallExpr); ok {
		cloneNested := clone.Args[offset].(*ast.CallExpr)

		return New(origNested, cloneNested, pass, handler)
	}

	return nil, false
}

func newWithTransformMatcher(fun ast.Expr, nested *Matcher, pass *analysis.Pass) *WithTransformMatcher {
	funcType := pass.TypesInfo.TypeOf(fun)
	if sig, ok := funcType.(*gotypes.Signature); ok && sig.Results().Len() > 0 {
		funcType = sig.Results().At(0).Type()
	}
	return &WithTransformMatcher{
		funcType: funcType,
		nested:   nested,
	}
}

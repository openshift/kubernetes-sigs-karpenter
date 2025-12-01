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

	"golang.org/x/tools/go/analysis"

	"github.com/nunnatsa/ginkgolinter/internal/gomegahandler"
)

const ( // gomega matchers
	beEmpty        = "BeEmpty"
	beEquivalentTo = "BeEquivalentTo"
	beFalse        = "BeFalse"
	beIdenticalTo  = "BeIdenticalTo"
	beNil          = "BeNil"
	beNumerically  = "BeNumerically"
	beTrue         = "BeTrue"
	beZero         = "BeZero"
	equal          = "Equal"
	haveLen        = "HaveLen"
	haveValue      = "HaveValue"
	and            = "And"
	or             = "Or"
	withTransform  = "WithTransform"
	matchError     = "MatchError"
	haveOccurred   = "HaveOccurred"
	succeed        = "Succeed"
)

type Matcher struct {
	funcName      string
	Orig          *ast.CallExpr
	Clone         *ast.CallExpr
	info          Info
	reverseLogic  bool
	handler       gomegahandler.Handler
	hasNotMatcher bool // true if the matcher is wrapped with a "Not" matcher
}

func New(origMatcher, matcherClone *ast.CallExpr, pass *analysis.Pass, handler gomegahandler.Handler) (*Matcher, bool) {
	reverse := false
	hasNotMatcher := false

	var assertFuncName string
	for {
		info, ok := handler.GetGomegaBasicInfo(origMatcher)
		if !ok {
			return nil, false
		}

		if info.MethodName != "Not" {
			assertFuncName = info.MethodName
			break
		}

		hasNotMatcher = true
		reverse = !reverse
		origMatcher, ok = origMatcher.Args[0].(*ast.CallExpr)
		if !ok {
			return nil, false
		}
		matcherClone = matcherClone.Args[0].(*ast.CallExpr)
	}

	return &Matcher{
		funcName:      assertFuncName,
		Orig:          origMatcher,
		Clone:         matcherClone,
		info:          getMatcherInfo(origMatcher, matcherClone, assertFuncName, pass, handler),
		reverseLogic:  reverse,
		hasNotMatcher: hasNotMatcher,
		handler:       handler,
	}, true
}

func (m *Matcher) ShouldReverseLogic() bool {
	return m.reverseLogic
}

func (m *Matcher) HasNotMatcher() bool {
	return m.hasNotMatcher
}

func (m *Matcher) GetMatcherInfo() Info {
	return m.info
}

func (m *Matcher) ReplaceMatcherFuncName(name string) {
	m.handler.ReplaceFunction(m.Clone, ast.NewIdent(name))
}

func (m *Matcher) ReplaceMatcherArgs(newArgs []ast.Expr) {
	m.Clone.Args = newArgs
}

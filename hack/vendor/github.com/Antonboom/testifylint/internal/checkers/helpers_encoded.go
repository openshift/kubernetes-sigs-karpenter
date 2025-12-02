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
	"go/token"
	"regexp"

	"golang.org/x/tools/go/analysis"

	"github.com/Antonboom/testifylint/internal/analysisutil"
)

var (
	wordsRe = regexp.MustCompile(`[A-Z]+(?:[a-z]*|$)|[a-z]+`) // NOTE(a.telyshev): ChatGPT.

	jsonIdentRe = regexp.MustCompile(`json|JSON|Json`)
	yamlWordRe  = regexp.MustCompile(`yaml|YAML|Yaml|^(yml|YML|Yml)$`)
)

func isJSONStyleExpr(pass *analysis.Pass, e ast.Expr) bool {
	if isIdentNamedAfterPattern(jsonIdentRe, e) {
		return hasBytesType(pass, e) || hasStringType(pass, e)
	}

	if t, ok := pass.TypesInfo.Types[e]; ok && t.Value != nil {
		return analysisutil.IsJSONLike(t.Value.String())
	}

	if bl, ok := e.(*ast.BasicLit); ok {
		return bl.Kind == token.STRING && analysisutil.IsJSONLike(bl.Value)
	}

	if args, ok := isFmtSprintfCall(pass, e); ok {
		return isJSONStyleExpr(pass, args[0])
	}

	return false
}

func isYAMLStyleExpr(pass *analysis.Pass, e ast.Expr) bool {
	id, ok := e.(*ast.Ident)
	return ok && (hasBytesType(pass, e) || hasStringType(pass, e)) && hasWordAfterPattern(id.Name, yamlWordRe)
}

func hasWordAfterPattern(s string, re *regexp.Regexp) bool {
	for _, w := range splitIntoWords(s) {
		if re.MatchString(w) {
			return true
		}
	}
	return false
}

func splitIntoWords(s string) []string {
	return wordsRe.FindAllString(s, -1)
}

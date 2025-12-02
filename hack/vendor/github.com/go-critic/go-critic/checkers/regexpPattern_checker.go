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
	"go/constant"
	"regexp"
	"strings"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "regexpPattern"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects suspicious regexp patterns"
	info.Before = "regexp.MustCompile(`google.com|yandex.ru`)"
	info.After = "regexp.MustCompile(`google\\.com|yandex\\.ru`)"

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		domains := []string{
			"com",
			"org",
			"info",
			"net",
			"ru",
			"de",
		}

		allDomains := strings.Join(domains, "|")
		domainRE := regexp.MustCompile(`[^\\]\.(` + allDomains + `)\b`)
		return astwalk.WalkerForExpr(&regexpPatternChecker{
			ctx:      ctx,
			domainRE: domainRE,
		}), nil
	})
}

type regexpPatternChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext

	domainRE *regexp.Regexp
}

func (c *regexpPatternChecker) VisitExpr(x ast.Expr) {
	call, ok := x.(*ast.CallExpr)
	if !ok {
		return
	}

	switch qualifiedName(call.Fun) {
	case "regexp.Compile", "regexp.CompilePOSIX", "regexp.MustCompile", "regexp.MustCompilePosix":
		cv := c.ctx.TypesInfo.Types[call.Args[0]].Value
		if cv == nil || cv.Kind() != constant.String {
			return
		}
		s := constant.StringVal(cv)
		if m := c.domainRE.FindStringSubmatch(s); m != nil {
			c.warnDomain(call.Args[0], m[1])
		}
	}
}

func (c *regexpPatternChecker) warnDomain(cause ast.Expr, domain string) {
	c.ctx.Warn(cause, "'.%s' should probably be '\\.%s'", domain, domain)
}

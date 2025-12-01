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

package rules

import (
	"go/ast"

	"github.com/securego/gosec/v2"
	"github.com/securego/gosec/v2/issue"
)

type httpServeWithoutTimeouts struct {
	issue.MetaData
	pkg   string
	calls []string
}

func (r *httpServeWithoutTimeouts) ID() string {
	return r.MetaData.ID
}

func (r *httpServeWithoutTimeouts) Match(n ast.Node, c *gosec.Context) (gi *issue.Issue, err error) {
	if _, matches := gosec.MatchCallByPackage(n, c, r.pkg, r.calls...); matches {
		return c.NewIssue(n, r.ID(), r.What, r.Severity, r.Confidence), nil
	}
	return nil, nil
}

// NewHTTPServeWithoutTimeouts detects use of net/http serve functions that have no support for setting timeouts.
func NewHTTPServeWithoutTimeouts(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	return &httpServeWithoutTimeouts{
		pkg:   "net/http",
		calls: []string{"ListenAndServe", "ListenAndServeTLS", "Serve", "ServeTLS"},
		MetaData: issue.MetaData{
			ID:         id,
			What:       "Use of net/http serve function that has no support for setting timeouts",
			Severity:   issue.Medium,
			Confidence: issue.High,
		},
	}, []ast.Node{(*ast.CallExpr)(nil)}
}

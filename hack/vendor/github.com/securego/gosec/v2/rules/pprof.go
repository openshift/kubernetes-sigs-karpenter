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

type pprofCheck struct {
	issue.MetaData
	importPath string
	importName string
}

// ID returns the ID of the check
func (p *pprofCheck) ID() string {
	return p.MetaData.ID
}

// Match checks for pprof imports
func (p *pprofCheck) Match(n ast.Node, c *gosec.Context) (*issue.Issue, error) {
	if node, ok := n.(*ast.ImportSpec); ok {
		if p.importPath == unquote(node.Path.Value) && node.Name != nil && p.importName == node.Name.Name {
			return c.NewIssue(node, p.ID(), p.What, p.Severity, p.Confidence), nil
		}
	}
	return nil, nil
}

// NewPprofCheck detects when the profiling endpoint is automatically exposed
func NewPprofCheck(id string, _ gosec.Config) (gosec.Rule, []ast.Node) {
	return &pprofCheck{
		MetaData: issue.MetaData{
			ID:         id,
			Severity:   issue.High,
			Confidence: issue.High,
			What:       "Profiling endpoint is automatically exposed on /debug/pprof",
		},
		importPath: "net/http/pprof",
		importName: "_",
	}, []ast.Node{(*ast.ImportSpec)(nil)}
}

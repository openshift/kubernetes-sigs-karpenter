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

package goanalysis

import (
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/result"
)

type Issue struct {
	*result.Issue
	Pass *analysis.Pass
}

func NewIssue(issue *result.Issue, pass *analysis.Pass) *Issue {
	return &Issue{
		Issue: issue,
		Pass:  pass,
	}
}

type EncodingIssue struct {
	FromLinter           string
	Text                 string
	Severity             string
	Pos                  token.Position
	LineRange            *result.Range
	SuggestedFixes       []analysis.SuggestedFix
	ExpectNoLint         bool
	ExpectedNoLintLinter string
}

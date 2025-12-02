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

package processors

import (
	"path/filepath"

	"github.com/golangci/golangci-lint/v2/pkg/logutils"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

var _ Processor = (*InvalidIssue)(nil)

// InvalidIssue filters invalid reports.
//   - non-go files (except `go.mod`)
//   - reports without file path
type InvalidIssue struct {
	log logutils.Log
}

func NewInvalidIssue(log logutils.Log) *InvalidIssue {
	return &InvalidIssue{log: log}
}

func (InvalidIssue) Name() string {
	return "invalid_issue"
}

func (p InvalidIssue) Process(issues []*result.Issue) ([]*result.Issue, error) {
	tcIssues := filterIssuesUnsafe(issues, func(issue *result.Issue) bool {
		return issue.FromLinter == typeCheckName
	})

	if len(tcIssues) > 0 {
		return tcIssues, nil
	}

	return filterIssuesErr(issues, p.shouldPassIssue)
}

func (InvalidIssue) Finish() {}

func (p InvalidIssue) shouldPassIssue(issue *result.Issue) (bool, error) {
	if issue.FilePath() == "" {
		p.log.Warnf("no file path for the issue: probably a bug inside the linter %q: %#v", issue.FromLinter, issue)

		return false, nil
	}

	if filepath.Base(issue.FilePath()) == "go.mod" {
		return true, nil
	}

	if !isGoFile(issue.FilePath()) {
		p.log.Infof("issue related to file %s is skipped", issue.FilePath())

		return false, nil
	}

	return true, nil
}

func isGoFile(name string) bool {
	return filepath.Ext(name) == ".go"
}

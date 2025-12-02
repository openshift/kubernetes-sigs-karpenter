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

var _ Processor = (*PathAbsoluter)(nil)

// PathAbsoluter ensures that representation of path are absolute.
type PathAbsoluter struct {
	log logutils.Log
}

func NewPathAbsoluter(log logutils.Log) *PathAbsoluter {
	return &PathAbsoluter{log: log.Child(logutils.DebugKeyPathAbsoluter)}
}

func (*PathAbsoluter) Name() string {
	return "path_absoluter"
}

func (p *PathAbsoluter) Process(issues []*result.Issue) ([]*result.Issue, error) {
	return transformIssues(issues, func(issue *result.Issue) *result.Issue {
		if filepath.IsAbs(issue.FilePath()) {
			return issue
		}

		absPath, err := filepath.Abs(issue.FilePath())
		if err != nil {
			p.log.Warnf("failed to get absolute path for %q: %v", issue.FilePath(), err)
			return nil
		}

		newIssue := issue
		newIssue.Pos.Filename = absPath

		return newIssue
	}), nil
}

func (*PathAbsoluter) Finish() {}

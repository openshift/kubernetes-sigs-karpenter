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
	"fmt"
	"strings"

	"github.com/golangci/golangci-lint/v2/pkg/fsutils"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

var _ Processor = (*PathShortener)(nil)

// PathShortener modifies text of the reports to reduce file path inside the text.
// It uses the rooted path name corresponding to the current directory (`wd`).
type PathShortener struct {
	wd string
}

func NewPathShortener() *PathShortener {
	wd, err := fsutils.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Can't get working dir: %s", err))
	}

	return &PathShortener{wd: wd}
}

func (PathShortener) Name() string {
	return "path_shortener"
}

func (p PathShortener) Process(issues []*result.Issue) ([]*result.Issue, error) {
	return transformIssues(issues, func(issue *result.Issue) *result.Issue {
		newIssue := issue
		newIssue.Text = strings.ReplaceAll(newIssue.Text, p.wd+"/", "")
		newIssue.Text = strings.ReplaceAll(newIssue.Text, p.wd, "")
		return newIssue
	}), nil
}

func (PathShortener) Finish() {}

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
	"path/filepath"

	"github.com/golangci/golangci-lint/v2/pkg/fsutils"
	"github.com/golangci/golangci-lint/v2/pkg/logutils"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

var _ Processor = (*PathRelativity)(nil)

// PathRelativity computes [result.Issue.RelativePath] and [result.Issue.WorkingDirectoryRelativePath],
// based on the base path.
type PathRelativity struct {
	log              logutils.Log
	basePath         string
	workingDirectory string
}

func NewPathRelativity(log logutils.Log, basePath string) (*PathRelativity, error) {
	wd, err := fsutils.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting working directory: %w", err)
	}

	return &PathRelativity{
		log:              log.Child(logutils.DebugKeyPathRelativity),
		basePath:         basePath,
		workingDirectory: wd,
	}, nil
}

func (*PathRelativity) Name() string {
	return "path_relativity"
}

func (p *PathRelativity) Process(issues []*result.Issue) ([]*result.Issue, error) {
	return transformIssues(issues, func(issue *result.Issue) *result.Issue {
		newIssue := *issue

		var err error
		newIssue.RelativePath, err = filepath.Rel(p.basePath, issue.FilePath())
		if err != nil {
			p.log.Warnf("Getting relative path (basepath): %v", err)
			return nil
		}

		newIssue.WorkingDirectoryRelativePath, err = filepath.Rel(p.workingDirectory, issue.FilePath())
		if err != nil {
			p.log.Warnf("Getting relative path (wd): %v", err)
			return nil
		}

		return &newIssue
	}), nil
}

func (*PathRelativity) Finish() {}

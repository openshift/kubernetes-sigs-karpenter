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

	"github.com/golangci/golangci-lint/v2/pkg/result"
)

func filterIssues(issues []*result.Issue, filter func(issue *result.Issue) bool) []*result.Issue {
	retIssues := make([]*result.Issue, 0, len(issues))
	for _, issue := range issues {
		if issue.FromLinter == typeCheckName {
			// don't hide typechecking errors in generated files: users expect to see why the project isn't compiling
			retIssues = append(retIssues, issue)
			continue
		}

		if filter(issue) {
			retIssues = append(retIssues, issue)
		}
	}

	return retIssues
}

func filterIssuesUnsafe(issues []*result.Issue, filter func(issue *result.Issue) bool) []*result.Issue {
	retIssues := make([]*result.Issue, 0, len(issues))
	for _, issue := range issues {
		if filter(issue) {
			retIssues = append(retIssues, issue)
		}
	}

	return retIssues
}

func filterIssuesErr(issues []*result.Issue, filter func(issue *result.Issue) (bool, error)) ([]*result.Issue, error) {
	retIssues := make([]*result.Issue, 0, len(issues))
	for _, issue := range issues {
		if issue.FromLinter == typeCheckName {
			// don't hide typechecking errors in generated files: users expect to see why the project isn't compiling
			retIssues = append(retIssues, issue)
			continue
		}

		ok, err := filter(issue)
		if err != nil {
			return nil, fmt.Errorf("can't filter issue %#v: %w", issue, err)
		}

		if ok {
			retIssues = append(retIssues, issue)
		}
	}

	return retIssues, nil
}

func transformIssues(issues []*result.Issue, transform func(issue *result.Issue) *result.Issue) []*result.Issue {
	retIssues := make([]*result.Issue, 0, len(issues))
	for _, issue := range issues {
		newIssue := transform(issue)
		if newIssue != nil {
			retIssues = append(retIssues, newIssue)
		}
	}

	return retIssues
}

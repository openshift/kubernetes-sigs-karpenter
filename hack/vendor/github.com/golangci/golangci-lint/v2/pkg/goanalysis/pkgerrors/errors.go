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

package pkgerrors

import (
	"errors"
	"fmt"

	"golang.org/x/tools/go/packages"

	"github.com/golangci/golangci-lint/v2/pkg/lint/linter"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

type IllTypedError struct {
	Pkg *packages.Package
}

func (e *IllTypedError) Error() string {
	return fmt.Sprintf("IllTypedError: errors in package: %v", e.Pkg.Errors)
}

func BuildIssuesFromIllTypedError(errs []error, lintCtx *linter.Context) ([]*result.Issue, error) {
	var issues []*result.Issue

	uniqReportedIssues := map[string]bool{}

	var other error

	for _, err := range errs {
		var ill *IllTypedError
		if !errors.As(err, &ill) {
			if other == nil {
				other = err
			}
			continue
		}

		for _, err := range extractErrors(ill.Pkg) {
			issue, perr := parseError(err)
			if perr != nil { // failed to parse
				if uniqReportedIssues[err.Msg] {
					continue
				}

				uniqReportedIssues[err.Msg] = true
				lintCtx.Log.Errorf("typechecking error: %s", err.Msg)
			} else {
				key := fmt.Sprintf("%s.%d.%d.%s", issue.FilePath(), issue.Line(), issue.Column(), issue.Text)
				if uniqReportedIssues[key] {
					continue
				}

				uniqReportedIssues[key] = true

				issue.Pkg = ill.Pkg // to save to cache later
				issues = append(issues, issue)
			}
		}
	}

	if len(issues) == 0 && other != nil {
		return nil, other
	}

	return issues, nil
}

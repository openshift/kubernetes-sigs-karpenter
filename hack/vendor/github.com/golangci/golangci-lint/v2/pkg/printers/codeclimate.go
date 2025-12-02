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

package printers

import (
	"encoding/json"
	"io"

	"github.com/golangci/golangci-lint/v2/pkg/logutils"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

const defaultCodeClimateSeverity = "critical"

// CodeClimate prints issues in the Code Climate format.
// https://github.com/codeclimate/platform/blob/HEAD/spec/analyzers/SPEC.md
type CodeClimate struct {
	log       logutils.Log
	w         io.Writer
	sanitizer severitySanitizer
}

func NewCodeClimate(log logutils.Log, w io.Writer) *CodeClimate {
	return &CodeClimate{
		log: log.Child(logutils.DebugKeyCodeClimatePrinter),
		w:   w,
		sanitizer: severitySanitizer{
			// https://github.com/codeclimate/platform/blob/HEAD/spec/analyzers/SPEC.md#data-types
			allowedSeverities: []string{"info", "minor", "major", defaultCodeClimateSeverity, "blocker"},
			defaultSeverity:   defaultCodeClimateSeverity,
		},
	}
}

func (p *CodeClimate) Print(issues []*result.Issue) error {
	ccIssues := make([]codeClimateIssue, 0, len(issues))

	for _, issue := range issues {
		ccIssue := codeClimateIssue{
			Description: issue.Description(),
			CheckName:   issue.FromLinter,
			Severity:    p.sanitizer.Sanitize(issue.Severity),
			Fingerprint: issue.Fingerprint(),
		}

		ccIssue.Location.Path = issue.Pos.Filename
		ccIssue.Location.Lines.Begin = issue.Pos.Line

		ccIssues = append(ccIssues, ccIssue)
	}

	err := p.sanitizer.Err()
	if err != nil {
		p.log.Infof("%v", err)
	}

	return json.NewEncoder(p.w).Encode(ccIssues)
}

// codeClimateIssue is a subset of the Code Climate spec.
// https://github.com/codeclimate/platform/blob/HEAD/spec/analyzers/SPEC.md#data-types
// It is just enough to support GitLab CI Code Quality.
// https://docs.gitlab.com/ee/ci/testing/code_quality.html#code-quality-report-format
type codeClimateIssue struct {
	Description string `json:"description"`
	CheckName   string `json:"check_name"`
	Severity    string `json:"severity,omitempty"`
	Fingerprint string `json:"fingerprint"`
	Location    struct {
		Path  string `json:"path"`
		Lines struct {
			Begin int `json:"begin"`
		} `json:"lines"`
	} `json:"location"`
}

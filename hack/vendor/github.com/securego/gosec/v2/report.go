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

package gosec

import (
	"github.com/securego/gosec/v2/issue"
)

// ReportInfo this is report information
type ReportInfo struct {
	Errors       map[string][]Error `json:"Golang errors"`
	Issues       []*issue.Issue
	Stats        *Metrics
	GosecVersion string
}

// NewReportInfo instantiate a ReportInfo
func NewReportInfo(issues []*issue.Issue, metrics *Metrics, errors map[string][]Error) *ReportInfo {
	return &ReportInfo{
		Errors: errors,
		Issues: issues,
		Stats:  metrics,
	}
}

// WithVersion defines the version of gosec used to generate the report
func (r *ReportInfo) WithVersion(version string) *ReportInfo {
	r.GosecVersion = version
	return r
}

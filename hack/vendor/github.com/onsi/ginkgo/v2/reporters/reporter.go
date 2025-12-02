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

package reporters

import (
	"github.com/onsi/ginkgo/v2/types"
)

type Reporter interface {
	SuiteWillBegin(report types.Report)
	WillRun(report types.SpecReport)
	DidRun(report types.SpecReport)
	SuiteDidEnd(report types.Report)

	//Timeline emission
	EmitFailure(state types.SpecState, failure types.Failure)
	EmitProgressReport(progressReport types.ProgressReport)
	EmitReportEntry(entry types.ReportEntry)
	EmitSpecEvent(event types.SpecEvent)
}

type NoopReporter struct{}

func (n NoopReporter) SuiteWillBegin(report types.Report)                       {}
func (n NoopReporter) WillRun(report types.SpecReport)                          {}
func (n NoopReporter) DidRun(report types.SpecReport)                           {}
func (n NoopReporter) SuiteDidEnd(report types.Report)                          {}
func (n NoopReporter) EmitFailure(state types.SpecState, failure types.Failure) {}
func (n NoopReporter) EmitProgressReport(progressReport types.ProgressReport)   {}
func (n NoopReporter) EmitReportEntry(entry types.ReportEntry)                  {}
func (n NoopReporter) EmitSpecEvent(event types.SpecEvent)                      {}

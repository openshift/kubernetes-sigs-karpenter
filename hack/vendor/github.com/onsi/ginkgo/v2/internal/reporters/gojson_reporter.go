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

type GoJSONReporter struct {
	ev *GoJSONEventWriter
}

type specSystemExtractFn func (spec types.SpecReport) string

func NewGoJSONReporter(enc encoder, errFn specSystemExtractFn, outFn specSystemExtractFn) *GoJSONReporter {
	return &GoJSONReporter{
		ev: NewGoJSONEventWriter(enc, errFn, outFn),
	}
}

func (r *GoJSONReporter) Write(originalReport types.Report) error {
	// suite start events
	report := newReport(originalReport)
	err := report.Fill()
	if err != nil {
		return err
	}
	r.ev.WriteSuiteStart(report)
	for _, originalSpecReport := range originalReport.SpecReports {
		specReport := newSpecReport(originalSpecReport)
		err := specReport.Fill()
		if err != nil {
			return err
		}
		if specReport.o.LeafNodeType == types.NodeTypeIt {
			// handle any It leaf node as a spec
			r.ev.WriteSpecStart(report, specReport)
			r.ev.WriteSpecOut(report, specReport)
			r.ev.WriteSpecResult(report, specReport)
		} else {
			// handle any other leaf node as generic output
			r.ev.WriteSpecOut(report, specReport)
		}
	}
	r.ev.WriteSuiteResult(report)
	return nil
}

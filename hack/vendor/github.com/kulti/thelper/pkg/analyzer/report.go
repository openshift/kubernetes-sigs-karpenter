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

package analyzer

import (
	"go/token"

	"golang.org/x/tools/go/analysis"
)

type reports struct {
	reports  []report
	filter   map[token.Pos]struct{}
	nofilter map[token.Pos]struct{}
}

type report struct {
	pos    token.Pos
	format string
	args   []interface{}
}

func (rr *reports) Reportf(pos token.Pos, format string, args ...interface{}) {
	rr.reports = append(rr.reports, report{
		pos:    pos,
		format: format,
		args:   args,
	})
}

func (rr *reports) Filter(pos token.Pos) {
	if pos.IsValid() {
		if rr.filter == nil {
			rr.filter = make(map[token.Pos]struct{})
		}

		rr.filter[pos] = struct{}{}
	}
}

func (rr *reports) NoFilter(pos token.Pos) {
	if pos.IsValid() {
		if rr.nofilter == nil {
			rr.nofilter = make(map[token.Pos]struct{})
		}

		rr.nofilter[pos] = struct{}{}
	}
}

func (rr *reports) Flush(pass *analysis.Pass) {
	for _, r := range rr.reports {
		if _, ok := rr.filter[r.pos]; ok {
			if _, ok := rr.nofilter[r.pos]; !ok {
				continue
			}
		}

		pass.Reportf(r.pos, r.format, r.args...)
	}
}

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

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package traces

import (
	"golang.org/x/vuln/internal/govulncheck"
)

// Compact returns a summarization of finding.Trace. The first
// returned element is the vulnerable symbol and the last element
// is the exit point of the user module. There can also be two
// elements in between, if applicable, which are the two elements
// preceding the user module exit point.
func Compact(finding *govulncheck.Finding) []*govulncheck.Frame {
	if len(finding.Trace) < 1 {
		return nil
	}
	iTop := len(finding.Trace) - 1
	topModule := finding.Trace[iTop].Module
	// search for the exit point of the top module
	for i, frame := range finding.Trace {
		if frame.Module == topModule {
			iTop = i
			break
		}
	}

	if iTop == 0 {
		// all in one module, reset to the end
		iTop = len(finding.Trace) - 1
	}

	compact := []*govulncheck.Frame{finding.Trace[0]}
	if iTop > 1 {
		if iTop > 2 {
			compact = append(compact, finding.Trace[iTop-2])
		}
		compact = append(compact, finding.Trace[iTop-1])
	}
	if iTop > 0 {
		compact = append(compact, finding.Trace[iTop])
	}
	return compact
}

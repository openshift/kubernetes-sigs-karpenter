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

package printer

// Golangci-lint: altered version of plumbing.go

import (
	"sort"

	"github.com/golangci/dupl/syntax"
)

type Clone clone

func (c Clone) Filename() string {
	return c.filename
}

func (c Clone) LineStart() int {
	return c.lineStart
}

func (c Clone) LineEnd() int {
	return c.lineEnd
}

type Issue struct {
	From, To Clone
}

type Issuer struct {
	ReadFile
}

func NewIssuer(fread ReadFile) *Issuer {
	return &Issuer{fread}
}

func (p *Issuer) MakeIssues(dups [][]*syntax.Node) ([]Issue, error) {
	clones, err := prepareClonesInfo(p.ReadFile, dups)
	if err != nil {
		return nil, err
	}

	sort.Sort(byNameAndLine(clones))

	var issues []Issue

	for i, cl := range clones {
		nextCl := clones[(i+1)%len(clones)]
		issues = append(issues, Issue{
			From: Clone(cl),
			To:   Clone(nextCl),
		})
	}

	return issues, nil
}

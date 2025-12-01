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

import (
	"fmt"
	"io"
	"sort"

	"github.com/golangci/dupl/syntax"
)

type plumbing struct {
	w io.Writer
	ReadFile
}

func NewPlumbing(w io.Writer, fread ReadFile) Printer {
	return &plumbing{w, fread}
}

func (p *plumbing) PrintHeader() error { return nil }

func (p *plumbing) PrintClones(dups [][]*syntax.Node) error {
	clones, err := prepareClonesInfo(p.ReadFile, dups)
	if err != nil {
		return err
	}
	sort.Sort(byNameAndLine(clones))
	for i, cl := range clones {
		nextCl := clones[(i+1)%len(clones)]
		fmt.Fprintf(p.w, "%s:%d-%d: duplicate of %s:%d-%d\n", cl.filename, cl.lineStart, cl.lineEnd,
			nextCl.filename, nextCl.lineStart, nextCl.lineEnd)
	}
	return nil
}

func (p *plumbing) PrintFooter() error { return nil }

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

// Package lib Golangci-lint: altered version of main.go
package lib

import (
	"os"
	"sort"

	"github.com/golangci/dupl/job"
	"github.com/golangci/dupl/printer"
	"github.com/golangci/dupl/syntax"
)

func Run(files []string, threshold int) ([]printer.Issue, error) {
	fchan := make(chan string, 1024)
	go func() {
		for _, f := range files {
			fchan <- f
		}
		close(fchan)
	}()
	schan := job.Parse(fchan)
	t, data, done := job.BuildTree(schan)
	<-done

	// finish stream
	t.Update(&syntax.Node{Type: -1})

	mchan := t.FindDuplOver(threshold)
	duplChan := make(chan syntax.Match)
	go func() {
		for m := range mchan {
			match := syntax.FindSyntaxUnits(*data, m, threshold)
			if len(match.Frags) > 0 {
				duplChan <- match
			}
		}
		close(duplChan)
	}()

	return makeIssues(duplChan)
}

func makeIssues(duplChan <-chan syntax.Match) ([]printer.Issue, error) {
	groups := make(map[string][][]*syntax.Node)
	for dupl := range duplChan {
		groups[dupl.Hash] = append(groups[dupl.Hash], dupl.Frags...)
	}
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	p := printer.NewIssuer(os.ReadFile)

	var issues []printer.Issue
	for _, k := range keys {
		uniq := unique(groups[k])
		if len(uniq) > 1 {
			i, err := p.MakeIssues(uniq)
			if err != nil {
				return nil, err
			}
			issues = append(issues, i...)
		}
	}

	return issues, nil
}

func unique(group [][]*syntax.Node) [][]*syntax.Node {
	fileMap := make(map[string]map[int]struct{})

	var newGroup [][]*syntax.Node
	for _, seq := range group {
		node := seq[0]
		file, ok := fileMap[node.Filename]
		if !ok {
			file = make(map[int]struct{})
			fileMap[node.Filename] = file
		}
		if _, ok := file[node.Pos]; !ok {
			file[node.Pos] = struct{}{}
			newGroup = append(newGroup, seq)
		}
	}
	return newGroup
}

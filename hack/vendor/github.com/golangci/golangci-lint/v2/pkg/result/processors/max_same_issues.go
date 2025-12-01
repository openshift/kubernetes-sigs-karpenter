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

package processors

import (
	"sort"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/logutils"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

var _ Processor = (*MaxSameIssues)(nil)

// MaxSameIssues limits the number of reports with the same text.
type MaxSameIssues struct {
	textCounter map[string]int
	limit       int
	log         logutils.Log
	cfg         *config.Config
}

func NewMaxSameIssues(limit int, log logutils.Log, cfg *config.Config) *MaxSameIssues {
	return &MaxSameIssues{
		textCounter: map[string]int{},
		limit:       limit,
		log:         log,
		cfg:         cfg,
	}
}

func (*MaxSameIssues) Name() string {
	return "max_same_issues"
}

func (p *MaxSameIssues) Process(issues []*result.Issue) ([]*result.Issue, error) {
	if p.limit <= 0 { // no limit
		return issues, nil
	}

	return filterIssuesUnsafe(issues, func(issue *result.Issue) bool {
		p.textCounter[issue.Text]++ // always inc for stat

		return p.textCounter[issue.Text] <= p.limit
	}), nil
}

func (p *MaxSameIssues) Finish() {
	walkStringToIntMapSortedByValue(p.textCounter, func(text string, count int) {
		if count > p.limit {
			p.log.Infof("%d/%d issues with text %q were hidden, use --max-same-issues",
				count-p.limit, count, text)
		}
	})
}

type kv struct {
	Key   string
	Value int
}

func walkStringToIntMapSortedByValue(m map[string]int, walk func(k string, v int)) {
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{
			Key:   k,
			Value: v,
		})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		walk(kv.Key, kv.Value)
	}
}

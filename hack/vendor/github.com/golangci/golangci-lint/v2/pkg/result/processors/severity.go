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
	"cmp"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/fsutils"
	"github.com/golangci/golangci-lint/v2/pkg/logutils"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

const severityFromLinter = "@linter"

var _ Processor = (*Severity)(nil)

// Severity modifies report severity.
// It uses the same `baseRule` structure as [ExcludeRules] processor.
//
// Warning: it doesn't use `path-prefix` option.
type Severity struct {
	name string

	log logutils.Log

	lines *fsutils.LineCache

	defaultSeverity string
	rules           []severityRule
}

func NewSeverity(log logutils.Log, lines *fsutils.LineCache, cfg *config.Severity) *Severity {
	p := &Severity{
		name:            "severity-rules",
		lines:           lines,
		log:             log,
		defaultSeverity: cfg.Default,
	}

	p.rules = parseRules(cfg.Rules, "", newSeverityRule)

	return p
}

func (p *Severity) Name() string { return p.name }

func (p *Severity) Process(issues []*result.Issue) ([]*result.Issue, error) {
	if len(p.rules) == 0 && p.defaultSeverity == "" {
		return issues, nil
	}

	return transformIssues(issues, p.transform), nil
}

func (*Severity) Finish() {}

func (p *Severity) transform(issue *result.Issue) *result.Issue {
	for _, rule := range p.rules {
		if rule.match(issue, p.lines, p.log) {
			if rule.severity == severityFromLinter || (rule.severity == "" && p.defaultSeverity == severityFromLinter) {
				return issue
			}

			issue.Severity = cmp.Or(rule.severity, p.defaultSeverity)

			return issue
		}
	}

	if p.defaultSeverity != severityFromLinter {
		issue.Severity = p.defaultSeverity
	}

	return issue
}

type severityRule struct {
	baseRule
	severity string
}

func newSeverityRule(rule *config.SeverityRule, prefix string) severityRule {
	return severityRule{
		baseRule: newBaseRule(&rule.BaseRule, prefix),
		severity: rule.Severity,
	}
}

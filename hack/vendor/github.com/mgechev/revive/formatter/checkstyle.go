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

package formatter

import (
	"bytes"
	"encoding/xml"
	plain "text/template"

	"github.com/mgechev/revive/lint"
)

// Checkstyle is an implementation of the Formatter interface
// which formats the errors to Checkstyle-like format.
type Checkstyle struct {
	Metadata lint.FormatterMetadata
}

// Name returns the name of the formatter.
func (*Checkstyle) Name() string {
	return "checkstyle"
}

type issue struct {
	Line       int
	Col        int
	What       string
	Confidence float64
	Severity   lint.Severity
	RuleName   string
}

// Format formats the failures gotten from the lint.
func (*Checkstyle) Format(failures <-chan lint.Failure, config lint.Config) (string, error) {
	issues := map[string][]issue{}
	for failure := range failures {
		buf := new(bytes.Buffer)
		xml.Escape(buf, []byte(failure.Failure))
		what := buf.String()
		iss := issue{
			Line:       failure.Position.Start.Line,
			Col:        failure.Position.Start.Column,
			What:       what,
			Confidence: failure.Confidence,
			Severity:   severity(config, failure),
			RuleName:   failure.RuleName,
		}
		fn := failure.Filename()
		if issues[fn] == nil {
			issues[fn] = []issue{}
		}
		issues[fn] = append(issues[fn], iss)
	}

	t, err := plain.New("revive").Parse(checkstyleTemplate)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, issues)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

const checkstyleTemplate = `<?xml version='1.0' encoding='UTF-8'?>
<checkstyle version="5.0">
{{- range $k, $v := . }}
    <file name="{{ $k }}">
      {{- range $i, $issue := $v }}
      <error line="{{ $issue.Line }}" column="{{ $issue.Col }}" message="{{ $issue.What }} (confidence {{ $issue.Confidence}})" severity="{{ $issue.Severity }}" source="revive/{{ $issue.RuleName }}"/>
      {{- end }}
    </file>
{{- end }}
</checkstyle>`

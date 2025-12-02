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

package printers

import (
	"encoding/xml"
	"fmt"
	"io"
	"maps"
	"slices"
	"strings"

	"github.com/go-xmlfmt/xmlfmt"

	"github.com/golangci/golangci-lint/v2/pkg/logutils"
	"github.com/golangci/golangci-lint/v2/pkg/result"
)

const defaultCheckstyleSeverity = "error"

// Checkstyle prints issues in the Checkstyle format.
// https://checkstyle.org/config.html
type Checkstyle struct {
	log       logutils.Log
	w         io.Writer
	sanitizer severitySanitizer
}

func NewCheckstyle(log logutils.Log, w io.Writer) *Checkstyle {
	return &Checkstyle{
		log: log.Child(logutils.DebugKeyCheckstylePrinter),
		w:   w,
		sanitizer: severitySanitizer{
			// https://checkstyle.org/config.html#Severity
			// https://checkstyle.org/property_types.html#SeverityLevel
			allowedSeverities: []string{"ignore", "info", "warning", defaultCheckstyleSeverity},
			defaultSeverity:   defaultCheckstyleSeverity,
		},
	}
}

func (p *Checkstyle) Print(issues []*result.Issue) error {
	out := checkstyleOutput{
		Version: "5.0",
	}

	files := map[string]*checkstyleFile{}

	for _, issue := range issues {
		file, ok := files[issue.FilePath()]
		if !ok {
			file = &checkstyleFile{
				Name: issue.FilePath(),
			}

			files[issue.FilePath()] = file
		}

		newError := &checkstyleError{
			Column:   issue.Column(),
			Line:     issue.Line(),
			Message:  issue.Text,
			Source:   issue.FromLinter,
			Severity: p.sanitizer.Sanitize(issue.Severity),
		}

		file.Errors = append(file.Errors, newError)
	}

	err := p.sanitizer.Err()
	if err != nil {
		p.log.Infof("%v", err)
	}

	out.Files = slices.SortedFunc(maps.Values(files), func(a *checkstyleFile, b *checkstyleFile) int {
		return strings.Compare(a.Name, b.Name)
	})

	data, err := xml.Marshal(&out)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(p.w, "%s%s\n", xml.Header, xmlfmt.FormatXML(string(data), "", "  "))
	if err != nil {
		return err
	}

	return nil
}

type checkstyleOutput struct {
	XMLName xml.Name          `xml:"checkstyle"`
	Version string            `xml:"version,attr"`
	Files   []*checkstyleFile `xml:"file"`
}

type checkstyleFile struct {
	Name   string             `xml:"name,attr"`
	Errors []*checkstyleError `xml:"error"`
}

type checkstyleError struct {
	Column   int    `xml:"column,attr"`
	Line     int    `xml:"line,attr"`
	Message  string `xml:"message,attr"`
	Severity string `xml:"severity,attr"`
	Source   string `xml:"source,attr"`
}

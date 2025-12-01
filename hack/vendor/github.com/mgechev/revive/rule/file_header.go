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

package rule

import (
	"fmt"
	"regexp"

	"github.com/mgechev/revive/lint"
)

// FileHeaderRule lints the header that each file should have.
type FileHeaderRule struct {
	header string
}

var (
	multiRegexp  = regexp.MustCompile(`^/\*`)
	singleRegexp = regexp.MustCompile("^//")
)

// Configure validates the rule configuration, and configures the rule accordingly.
//
// Configuration implements the [lint.ConfigurableRule] interface.
func (r *FileHeaderRule) Configure(arguments lint.Arguments) error {
	if len(arguments) < 1 {
		return nil
	}

	var ok bool
	r.header, ok = arguments[0].(string)
	if !ok {
		return fmt.Errorf(`invalid argument for "file-header" rule: argument should be a string, got %T`, arguments[0])
	}
	return nil
}

// Apply applies the rule to given file.
func (r *FileHeaderRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	if r.header == "" {
		return nil
	}

	failure := []lint.Failure{
		{
			Node:       file.AST,
			Confidence: 1,
			Failure:    "the file doesn't have an appropriate header",
		},
	}

	if len(file.AST.Comments) == 0 {
		return failure
	}

	g := file.AST.Comments[0]
	if g == nil {
		return failure
	}
	comment := ""
	for _, c := range g.List {
		text := c.Text
		if multiRegexp.MatchString(text) {
			text = text[2 : len(text)-2]
		} else if singleRegexp.MatchString(text) {
			text = text[2:]
		}
		comment += text
	}

	regex, err := regexp.Compile(r.header)
	if err != nil {
		return newInternalFailureError(err)
	}

	if !regex.MatchString(comment) {
		return failure
	}
	return nil
}

// Name returns the rule name.
func (*FileHeaderRule) Name() string {
	return "file-header"
}

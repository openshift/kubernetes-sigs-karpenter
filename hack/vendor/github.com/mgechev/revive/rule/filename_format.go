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
	"path/filepath"
	"regexp"
	"unicode"

	"github.com/mgechev/revive/lint"
)

// FilenameFormatRule lints source filenames according to a set of regular expressions given as arguments.
type FilenameFormatRule struct {
	format *regexp.Regexp
}

// Apply applies the rule to the given file.
func (r *FilenameFormatRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	filename := filepath.Base(file.Name)
	if r.format.MatchString(filename) {
		return nil
	}

	failureMsg := fmt.Sprintf("Filename %s is not of the format %s.%s", filename, r.format.String(), r.getMsgForNonASCIIChars(filename))
	return []lint.Failure{{
		Confidence: 1,
		Failure:    failureMsg,
		RuleName:   r.Name(),
		Node:       file.AST.Name,
	}}
}

func (*FilenameFormatRule) getMsgForNonASCIIChars(str string) string {
	result := ""
	for _, c := range str {
		if c <= unicode.MaxASCII {
			continue
		}

		result += fmt.Sprintf(" Non ASCII character %c (%U) found.", c, c)
	}

	return result
}

// Name returns the rule name.
func (*FilenameFormatRule) Name() string {
	return "filename-format"
}

var defaultFormat = regexp.MustCompile(`^[_A-Za-z0-9][_A-Za-z0-9-]*\.go$`)

// Configure validates the rule configuration, and configures the rule accordingly.
//
// Configuration implements the [lint.ConfigurableRule] interface.
func (r *FilenameFormatRule) Configure(arguments lint.Arguments) error {
	argsCount := len(arguments)
	if argsCount == 0 {
		r.format = defaultFormat
		return nil
	}

	if argsCount > 1 {
		return fmt.Errorf("rule %q expects only one argument, got %d %v", r.Name(), argsCount, arguments)
	}

	arg := arguments[0]
	str, ok := arg.(string)
	if !ok {
		return fmt.Errorf("rule %q expects a string argument, got %v of type %T", r.Name(), arg, arg)
	}

	format, err := regexp.Compile(str)
	if err != nil {
		return fmt.Errorf("rule %q expects a valid regexp argument, got error for %s: %w", r.Name(), str, err)
	}

	r.format = format

	return nil
}

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

package internal

import (
	"fmt"
	"strings"
	"unicode"
)

func formatExtraLeadingSpace(fullDirective string) string {
	return fmt.Sprintf("directive `%s` should not have more than one leading space", fullDirective)
}

func formatNotMachine(fullDirective string) string {
	expected := fullDirective[:2] + strings.TrimLeftFunc(fullDirective[2:], unicode.IsSpace)
	return fmt.Sprintf("directive `%s` should be written without leading space as `%s`",
		fullDirective, expected)
}

func formatNotSpecific(fullDirective, directiveWithOptionalLeadingSpace string) string {
	return fmt.Sprintf("directive `%s` should mention specific linter such as `%s:my-linter`",
		fullDirective, directiveWithOptionalLeadingSpace)
}

func formatParseError(fullDirective, directiveWithOptionalLeadingSpace string) string {
	return fmt.Sprintf("directive `%s` should match `%s[:<comma-separated-linters>] [// <explanation>]`",
		fullDirective,
		directiveWithOptionalLeadingSpace)
}

func formatNoExplanation(fullDirective, fullDirectiveWithoutExplanation string) string {
	return fmt.Sprintf("directive `%s` should provide explanation such as `%s // this is why`",
		fullDirective, fullDirectiveWithoutExplanation)
}

func formatUnusedCandidate(fullDirective, expectedLinter string) string {
	details := fmt.Sprintf("directive `%s` is unused", fullDirective)
	if expectedLinter != "" {
		details += fmt.Sprintf(" for linter %q", expectedLinter)
	}
	return details
}

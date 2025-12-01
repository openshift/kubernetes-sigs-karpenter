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
	"fmt"

	"github.com/mgechev/revive/lint"
)

// Default is an implementation of the Formatter interface
// which formats the errors to text.
type Default struct {
	Metadata lint.FormatterMetadata
}

// Name returns the name of the formatter.
func (*Default) Name() string {
	return "default"
}

// Format formats the failures gotten from the lint.
func (*Default) Format(failures <-chan lint.Failure, _ lint.Config) (string, error) {
	var buf bytes.Buffer
	prefix := ""
	for failure := range failures {
		fmt.Fprintf(&buf, "%s%v: %s", prefix, failure.Position.Start, failure.Failure)
		prefix = "\n"
	}
	return buf.String(), nil
}

func ruleDescriptionURL(ruleName string) string {
	return "https://revive.run/r#" + ruleName
}

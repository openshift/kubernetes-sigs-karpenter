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
	"fmt"
	"strings"

	"github.com/mgechev/revive/lint"
)

// Unix is an implementation of the Formatter interface
// which formats the errors to a simple line based error format
//
//	main.go:24:9: [errorf] should replace errors.New(fmt.Sprintf(...)) with fmt.Errorf(...)
type Unix struct {
	Metadata lint.FormatterMetadata
}

// Name returns the name of the formatter.
func (*Unix) Name() string {
	return "unix"
}

// Format formats the failures gotten from the lint.
func (*Unix) Format(failures <-chan lint.Failure, _ lint.Config) (string, error) {
	var sb strings.Builder
	for failure := range failures {
		sb.WriteString(fmt.Sprintf("%v: [%s] %s\n", failure.Position.Start, failure.RuleName, failure.Failure))
	}
	return sb.String(), nil
}

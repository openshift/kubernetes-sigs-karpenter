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

package gosec

import (
	"sort"
)

// Error is used when there are golang errors while parsing the AST
type Error struct {
	Line   int    `json:"line"`
	Column int    `json:"column"`
	Err    string `json:"error"`
}

// NewError creates Error object
func NewError(line, column int, err string) *Error {
	return &Error{
		Line:   line,
		Column: column,
		Err:    err,
	}
}

// sortErrors sorts the golang errors by line
func sortErrors(allErrors map[string][]Error) {
	for _, errors := range allErrors {
		sort.Slice(errors, func(i, j int) bool {
			if errors[i].Line == errors[j].Line {
				return errors[i].Column <= errors[j].Column
			}
			return errors[i].Line < errors[j].Line
		})
	}
}

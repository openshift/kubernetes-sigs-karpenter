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

package gomodguard

import (
	"fmt"
	"go/token"
)

// Issue represents the result of one error.
type Issue struct {
	FileName   string
	LineNumber int
	Position   token.Position
	Reason     string
}

// String returns the filename, line
// number and reason of a Issue.
func (r *Issue) String() string {
	return fmt.Sprintf("%s:%d:1 %s", r.FileName, r.LineNumber, r.Reason)
}

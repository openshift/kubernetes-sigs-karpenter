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

// Position support for go-toml

package toml

import (
	"fmt"
)

// Position of a document element within a TOML document.
//
// Line and Col are both 1-indexed positions for the element's line number and
// column number, respectively.  Values of zero or less will cause Invalid(),
// to return true.
type Position struct {
	Line int // line within the document
	Col  int // column within the line
}

// String representation of the position.
// Displays 1-indexed line and column numbers.
func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.Line, p.Col)
}

// Invalid returns whether or not the position is valid (i.e. with negative or
// null values)
func (p Position) Invalid() bool {
	return p.Line <= 0 || p.Col <= 0
}

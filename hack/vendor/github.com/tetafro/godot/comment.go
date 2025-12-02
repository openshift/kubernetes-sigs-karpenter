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

package godot

import "go/token"

// comment is an internal representation of AST comment entity with additional
// data attached. The latter is used for creating a full replacement for
// the line with issues.
type comment struct {
	lines []string       // unmodified lines from file
	text  string         // concatenated `lines` with special parts excluded
	start token.Position // position of the first symbol in comment
	decl  bool           // whether comment is a declaration comment
}

// position is a position inside a comment (might be multiline comment).
type position struct {
	line   int // starts at 1
	column int // starts at 1, byte count
}

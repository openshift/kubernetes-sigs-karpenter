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

package revgrep

// Issue contains metadata about an issue found.
type Issue struct {
	// File is the name of the file as it appeared from the patch.
	File string
	// LineNo is the line number of the file.
	LineNo int
	// ColNo is the column number or 0 if none could be parsed.
	ColNo int
	// HunkPos is position from file's first @@, for new files this will be the line number.
	// See also: https://developer.github.com/v3/pulls/comments/#create-a-comment
	HunkPos int
	// Issue text as it appeared from the tool.
	Issue string
	// Message is the issue without file name, line number and column number.
	Message string
}

// InputIssue represents issue found by some linter.
type InputIssue interface {
	FilePath() string
	Line() int
}

type simpleInputIssue struct {
	filePath   string
	lineNumber int
}

func (i simpleInputIssue) FilePath() string {
	return i.filePath
}

func (i simpleInputIssue) Line() int {
	return i.lineNumber
}

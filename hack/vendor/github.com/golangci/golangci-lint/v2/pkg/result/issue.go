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

package result

import (
	"crypto/md5" //nolint:gosec // for md5 hash
	"fmt"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/packages"
)

type Range struct {
	From, To int
}

type Issue struct {
	FromLinter string
	Text       string

	Severity string

	// Source lines of a code with the issue to show
	SourceLines []string

	// Pkg is needed for proper caching of linting results
	Pkg *packages.Package `json:"-"`

	Pos token.Position

	LineRange *Range `json:",omitempty"`

	// HunkPos is used only when golangci-lint is run over a diff
	HunkPos int `json:",omitempty"`

	// If we know how to fix the issue, we can provide replacement lines
	SuggestedFixes []analysis.SuggestedFix `json:",omitempty"`

	// If we are expecting a nolint (because this is from nolintlint), record the expected linter
	ExpectNoLint         bool
	ExpectedNoLintLinter string

	// Only for Diff processor needs.
	WorkingDirectoryRelativePath string `json:"-"`

	// Only for processors that need relative paths evaluation.
	RelativePath string `json:"-"`
}

func (i *Issue) FilePath() string {
	return i.Pos.Filename
}

func (i *Issue) Line() int {
	return i.Pos.Line
}

func (i *Issue) Column() int {
	return i.Pos.Column
}

func (i *Issue) GetLineRange() Range {
	if i.LineRange == nil {
		return Range{
			From: i.Line(),
			To:   i.Line(),
		}
	}

	if i.LineRange.From == 0 {
		return Range{
			From: i.Line(),
			To:   i.Line(),
		}
	}

	return *i.LineRange
}

func (i *Issue) Description() string {
	return fmt.Sprintf("%s: %s", i.FromLinter, i.Text)
}

func (i *Issue) Fingerprint() string {
	firstLine := ""
	if len(i.SourceLines) > 0 {
		firstLine = i.SourceLines[0]
	}

	hash := md5.New() //nolint:gosec // we don't need a strong hash here
	_, _ = fmt.Fprintf(hash, "%s%s%s", i.Pos.Filename, i.Text, firstLine)

	return fmt.Sprintf("%X", hash.Sum(nil))
}

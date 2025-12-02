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

package migrate

import (
	"github.com/golangci/golangci-lint/v2/pkg/commands/internal/migrate/versionone"
	"github.com/golangci/golangci-lint/v2/pkg/commands/internal/migrate/versiontwo"
)

func toIssues(old *versionone.Config) versiontwo.Issues {
	return versiontwo.Issues{
		MaxIssuesPerLinter: old.Issues.MaxIssuesPerLinter,
		MaxSameIssues:      old.Issues.MaxSameIssues,
		UniqByLine:         old.Issues.UniqByLine,
		DiffFromRevision:   old.Issues.DiffFromRevision,
		DiffFromMergeBase:  old.Issues.DiffFromMergeBase,
		DiffPatchFilePath:  old.Issues.DiffPatchFilePath,
		WholeFiles:         old.Issues.WholeFiles,
		Diff:               old.Issues.Diff,
		NeedFix:            old.Issues.NeedFix,
	}
}

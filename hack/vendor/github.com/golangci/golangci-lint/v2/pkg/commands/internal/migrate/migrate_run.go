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
	"github.com/golangci/golangci-lint/v2/pkg/commands/internal/migrate/ptr"
	"github.com/golangci/golangci-lint/v2/pkg/commands/internal/migrate/versionone"
	"github.com/golangci/golangci-lint/v2/pkg/commands/internal/migrate/versiontwo"
)

func toRun(old *versionone.Config) versiontwo.Run {
	var relativePathMode *string
	if ptr.Deref(old.Run.RelativePathMode) != "cfg" {
		// cfg is the new default.
		relativePathMode = old.Run.RelativePathMode
	}

	var concurrency *int
	if ptr.Deref(old.Run.Concurrency) != 0 {
		// 0 is the new default
		concurrency = old.Run.Concurrency
	}

	return versiontwo.Run{
		Timeout:               0, // Enforce new default.
		Concurrency:           concurrency,
		Go:                    old.Run.Go,
		RelativePathMode:      relativePathMode,
		BuildTags:             old.Run.BuildTags,
		ModulesDownloadMode:   old.Run.ModulesDownloadMode,
		ExitCodeIfIssuesFound: old.Run.ExitCodeIfIssuesFound,
		AnalyzeTests:          old.Run.AnalyzeTests,
		AllowParallelRunners:  old.Run.AllowParallelRunners,
		AllowSerialRunners:    old.Run.AllowSerialRunners,
	}
}

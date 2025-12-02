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

package wsl

import (
	"github.com/bombsimon/wsl/v5"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
)

func NewV5(settings *config.WSLv5Settings) *goanalysis.Linter {
	var conf *wsl.Configuration

	if settings != nil {
		checkSet, err := wsl.NewCheckSet(settings.Default, settings.Enable, settings.Disable)
		if err != nil {
			internal.LinterLogger.Fatalf("wsl: invalid check: %v", err)
		}

		conf = &wsl.Configuration{
			IncludeGenerated:  true, // force to true because golangci-lint already has a way to filter generated files.
			AllowFirstInBlock: settings.AllowFirstInBlock,
			AllowWholeBlock:   settings.AllowWholeBlock,
			BranchMaxLines:    settings.BranchMaxLines,
			CaseMaxLines:      settings.CaseMaxLines,
			Checks:            checkSet,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(wsl.NewAnalyzer(conf)).
		WithVersion(5). //nolint:mnd // It's the linter version.
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

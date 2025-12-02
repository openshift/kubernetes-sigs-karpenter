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

package perfsprint

import (
	"github.com/catenacyber/perfsprint/analyzer"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.PerfSprintSettings) *goanalysis.Linter {
	cfg := map[string]any{
		"fiximports": false,
	}

	if settings != nil {
		// NOTE: The option `ignore-tests` is not handled because it should be managed with `linters.exclusions.rules`

		cfg["integer-format"] = settings.IntegerFormat
		cfg["int-conversion"] = settings.IntConversion

		cfg["error-format"] = settings.ErrorFormat
		cfg["err-error"] = settings.ErrError
		cfg["errorf"] = settings.ErrorF

		cfg["string-format"] = settings.StringFormat
		cfg["sprintf1"] = settings.SprintF1
		cfg["strconcat"] = settings.StrConcat

		cfg["bool-format"] = settings.BoolFormat
		cfg["hex-format"] = settings.HexFormat

		cfg["concat-loop"] = settings.ConcatLoop
		cfg["loop-other-ops"] = settings.LoopOtherOps
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer.New()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

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

package gosmopolitan

import (
	"strings"

	"github.com/xen0n/gosmopolitan"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.GosmopolitanSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			"allowtimelocal":  settings.AllowTimeLocal,
			"escapehatches":   strings.Join(settings.EscapeHatches, ","),
			"watchforscripts": strings.Join(settings.WatchForScripts, ","),

			// Should be managed with `linters.exclusions.rules`.
			"lookattests": true,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(gosmopolitan.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

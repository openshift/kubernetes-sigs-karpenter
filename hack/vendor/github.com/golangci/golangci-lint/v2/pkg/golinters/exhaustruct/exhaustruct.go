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

package exhaustruct

import (
	exhaustruct "dev.gaijin.team/go/exhaustruct/v4/analyzer"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
)

func New(settings *config.ExhaustructSettings) *goanalysis.Linter {
	cfg := exhaustruct.Config{}
	if settings != nil {
		cfg.IncludeRx = settings.Include
		cfg.ExcludeRx = settings.Exclude
		cfg.AllowEmpty = settings.AllowEmpty
		cfg.AllowEmptyRx = settings.AllowEmptyRx
		cfg.AllowEmptyReturns = settings.AllowEmptyReturns
		cfg.AllowEmptyDeclarations = settings.AllowEmptyDeclarations
	}

	analyzer, err := exhaustruct.NewAnalyzer(cfg)
	if err != nil {
		internal.LinterLogger.Fatalf("exhaustruct configuration: %v", err)
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

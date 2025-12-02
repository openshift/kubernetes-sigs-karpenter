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

package usestdlibvars

import (
	"github.com/sashamelentyev/usestdlibvars/pkg/analyzer"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.UseStdlibVarsSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			analyzer.ConstantKindFlag:       settings.ConstantKind,
			analyzer.CryptoHashFlag:         settings.CryptoHash,
			analyzer.HTTPMethodFlag:         settings.HTTPMethod,
			analyzer.HTTPStatusCodeFlag:     settings.HTTPStatusCode,
			analyzer.OSDevNullFlag:          false, // Noop because the linter ignore it.
			analyzer.RPCDefaultPathFlag:     settings.DefaultRPCPath,
			analyzer.SQLIsolationLevelFlag:  settings.SQLIsolationLevel,
			analyzer.SyslogPriorityFlag:     false, // Noop because the linter ignore it.
			analyzer.TimeLayoutFlag:         settings.TimeLayout,
			analyzer.TimeMonthFlag:          settings.TimeMonth,
			analyzer.TimeWeekdayFlag:        settings.TimeWeekday,
			analyzer.TLSSignatureSchemeFlag: settings.TLSSignatureScheme,
			analyzer.TimeDateMonthFlag:      settings.TimeDateMonth,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer.New()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

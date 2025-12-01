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

package loggercheck

import (
	"github.com/timonwong/loggercheck"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.LoggerCheckSettings) *goanalysis.Linter {
	var opts []loggercheck.Option

	if settings != nil {
		var disable []string
		if !settings.Kitlog {
			disable = append(disable, "kitlog")
		}
		if !settings.Klog {
			disable = append(disable, "klog")
		}
		if !settings.Logr {
			disable = append(disable, "logr")
		}
		if !settings.Slog {
			disable = append(disable, "slog")
		}
		if !settings.Zap {
			disable = append(disable, "zap")
		}

		opts = []loggercheck.Option{
			loggercheck.WithDisable(disable),
			loggercheck.WithRequireStringKey(settings.RequireStringKey),
			loggercheck.WithRules(settings.Rules),
			loggercheck.WithNoPrintfLike(settings.NoPrintfLike),
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(loggercheck.NewAnalyzer(opts...)).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

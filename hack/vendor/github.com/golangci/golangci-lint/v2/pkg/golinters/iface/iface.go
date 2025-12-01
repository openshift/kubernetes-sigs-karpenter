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

package iface

import (
	"slices"

	"github.com/uudashr/iface/identical"
	"github.com/uudashr/iface/opaque"
	"github.com/uudashr/iface/unexported"
	"github.com/uudashr/iface/unused"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.IfaceSettings) *goanalysis.Linter {
	var conf map[string]map[string]any
	if settings != nil {
		conf = settings.Settings
	}

	return goanalysis.NewLinter(
		"iface",
		"Detect the incorrect use of interfaces, helping developers avoid interface pollution.",
		analyzersFromSettings(settings),
		conf,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}

func analyzersFromSettings(settings *config.IfaceSettings) []*analysis.Analyzer {
	allAnalyzers := map[string]*analysis.Analyzer{
		"identical":  identical.Analyzer,
		"unused":     unused.Analyzer,
		"opaque":     opaque.Analyzer,
		"unexported": unexported.Analyzer,
	}

	if settings == nil || len(settings.Enable) == 0 {
		// Default enable `identical` analyzer only
		return []*analysis.Analyzer{identical.Analyzer}
	}

	var analyzers []*analysis.Analyzer
	for _, name := range uniqueNames(settings.Enable) {
		if _, ok := allAnalyzers[name]; !ok {
			// skip unknown analyzer
			continue
		}

		analyzers = append(analyzers, allAnalyzers[name])
	}

	return analyzers
}

func uniqueNames(names []string) []string {
	slices.Sort(names)
	return slices.Compact(names)
}

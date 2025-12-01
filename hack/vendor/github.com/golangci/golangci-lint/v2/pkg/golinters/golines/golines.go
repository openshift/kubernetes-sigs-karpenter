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

package golines

import (
	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/goformatters"
	golinesbase "github.com/golangci/golangci-lint/v2/pkg/goformatters/golines"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
)

func New(settings *config.GoLinesSettings) *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(golinesbase.Name),
				"Checks if code is formatted, and fixes long lines",
				golinesbase.New(settings),
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

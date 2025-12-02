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

package funlen

import (
	"github.com/ultraware/funlen"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

type Config struct {
	lineLimit      int
	stmtLimit      int
	ignoreComments bool
}

func New(settings *config.FunlenSettings) *goanalysis.Linter {
	cfg := Config{}
	if settings != nil {
		cfg.lineLimit = settings.Lines
		cfg.stmtLimit = settings.Statements
		cfg.ignoreComments = settings.IgnoreComments
	}

	return goanalysis.
		NewLinterFromAnalyzer(funlen.NewAnalyzer(cfg.lineLimit, cfg.stmtLimit, cfg.ignoreComments)).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

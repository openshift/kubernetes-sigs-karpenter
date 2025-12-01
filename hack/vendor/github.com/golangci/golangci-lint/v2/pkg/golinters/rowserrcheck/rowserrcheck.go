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

package rowserrcheck

import (
	"github.com/jingyugao/rowserrcheck/passes/rowserr"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.RowsErrCheckSettings) *goanalysis.Linter {
	var pkgs []string

	if settings != nil {
		pkgs = settings.Packages
	}

	return goanalysis.
		NewLinterFromAnalyzer(rowserr.NewAnalyzer(pkgs...)).
		WithDesc("checks whether Rows.Err of rows is checked successfully").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

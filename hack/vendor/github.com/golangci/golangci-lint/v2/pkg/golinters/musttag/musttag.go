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

package musttag

import (
	"go-simpler.org/musttag"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.MustTagSettings) *goanalysis.Linter {
	var funcs []musttag.Func

	if settings != nil {
		for _, fn := range settings.Functions {
			funcs = append(funcs, musttag.Func{
				Name:   fn.Name,
				Tag:    fn.Tag,
				ArgPos: fn.ArgPos,
			})
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(musttag.New(funcs...)).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

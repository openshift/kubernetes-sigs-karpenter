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

package tagalign

import (
	"github.com/4meepo/tagalign"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.TagAlignSettings) *goanalysis.Linter {
	var options []tagalign.Option

	if settings != nil {
		options = append(options, tagalign.WithAlign(settings.Align))

		if settings.Sort || len(settings.Order) > 0 {
			options = append(options, tagalign.WithSort(settings.Order...))
		}

		// Strict style will be applied only if Align and Sort are enabled together.
		if settings.Strict && settings.Align && settings.Sort {
			options = append(options, tagalign.WithStrictStyle())
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(tagalign.NewAnalyzer(options...)).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

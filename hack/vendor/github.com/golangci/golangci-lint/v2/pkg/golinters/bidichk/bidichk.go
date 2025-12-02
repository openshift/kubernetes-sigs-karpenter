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

package bidichk

import (
	"strings"

	"github.com/breml/bidichk/pkg/bidichk"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.BiDiChkSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		var opts []string

		if settings.LeftToRightEmbedding {
			opts = append(opts, "LEFT-TO-RIGHT-EMBEDDING")
		}
		if settings.RightToLeftEmbedding {
			opts = append(opts, "RIGHT-TO-LEFT-EMBEDDING")
		}
		if settings.PopDirectionalFormatting {
			opts = append(opts, "POP-DIRECTIONAL-FORMATTING")
		}
		if settings.LeftToRightOverride {
			opts = append(opts, "LEFT-TO-RIGHT-OVERRIDE")
		}
		if settings.RightToLeftOverride {
			opts = append(opts, "RIGHT-TO-LEFT-OVERRIDE")
		}
		if settings.LeftToRightIsolate {
			opts = append(opts, "LEFT-TO-RIGHT-ISOLATE")
		}
		if settings.RightToLeftIsolate {
			opts = append(opts, "RIGHT-TO-LEFT-ISOLATE")
		}
		if settings.FirstStrongIsolate {
			opts = append(opts, "FIRST-STRONG-ISOLATE")
		}
		if settings.PopDirectionalIsolate {
			opts = append(opts, "POP-DIRECTIONAL-ISOLATE")
		}

		cfg = map[string]any{
			"disallowed-runes": strings.Join(opts, ","),
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(bidichk.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

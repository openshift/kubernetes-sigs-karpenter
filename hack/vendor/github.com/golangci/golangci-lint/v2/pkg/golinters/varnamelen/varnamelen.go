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

package varnamelen

import (
	"strconv"
	"strings"

	"github.com/blizzy78/varnamelen"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.VarnamelenSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			"checkReceiver":      strconv.FormatBool(settings.CheckReceiver),
			"checkReturn":        strconv.FormatBool(settings.CheckReturn),
			"checkTypeParam":     strconv.FormatBool(settings.CheckTypeParam),
			"ignoreNames":        strings.Join(settings.IgnoreNames, ","),
			"ignoreTypeAssertOk": strconv.FormatBool(settings.IgnoreTypeAssertOk),
			"ignoreMapIndexOk":   strconv.FormatBool(settings.IgnoreMapIndexOk),
			"ignoreChanRecvOk":   strconv.FormatBool(settings.IgnoreChanRecvOk),
			"ignoreDecls":        strings.Join(settings.IgnoreDecls, ","),
		}

		if settings.MaxDistance > 0 {
			cfg["maxDistance"] = strconv.Itoa(settings.MaxDistance)
		}

		if settings.MinNameLength > 0 {
			cfg["minNameLength"] = strconv.Itoa(settings.MinNameLength)
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(varnamelen.NewAnalyzer()).
		WithDesc("checks that the length of a variable's name matches its scope").
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

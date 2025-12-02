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

package decorder

import (
	"strings"

	"gitlab.com/bosi/decorder"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.DecorderSettings) *goanalysis.Linter {
	// disable all rules/checks by default
	cfg := map[string]any{
		"ignore-underscore-vars":        false,
		"disable-dec-num-check":         true,
		"disable-type-dec-num-check":    false,
		"disable-const-dec-num-check":   false,
		"disable-var-dec-num-check":     false,
		"disable-dec-order-check":       true,
		"disable-init-func-first-check": true,
	}

	if settings != nil {
		cfg["dec-order"] = strings.Join(settings.DecOrder, ",")
		cfg["ignore-underscore-vars"] = settings.IgnoreUnderscoreVars
		cfg["disable-dec-num-check"] = settings.DisableDecNumCheck
		cfg["disable-type-dec-num-check"] = settings.DisableTypeDecNumCheck
		cfg["disable-const-dec-num-check"] = settings.DisableConstDecNumCheck
		cfg["disable-var-dec-num-check"] = settings.DisableVarDecNumCheck
		cfg["disable-dec-order-check"] = settings.DisableDecOrderCheck
		cfg["disable-init-func-first-check"] = settings.DisableInitFuncFirstCheck
	}

	return goanalysis.
		NewLinterFromAnalyzer(decorder.Analyzer).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeSyntax)
}

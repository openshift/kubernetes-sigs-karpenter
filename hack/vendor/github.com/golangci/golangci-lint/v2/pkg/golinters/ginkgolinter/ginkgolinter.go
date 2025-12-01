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

package ginkgolinter

import (
	"github.com/nunnatsa/ginkgolinter"
	glconfig "github.com/nunnatsa/ginkgolinter/config"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.GinkgoLinterSettings) *goanalysis.Linter {
	cfg := &glconfig.Config{}

	if settings != nil {
		cfg = &glconfig.Config{
			SuppressLen:               settings.SuppressLenAssertion,
			SuppressNil:               settings.SuppressNilAssertion,
			SuppressErr:               settings.SuppressErrAssertion,
			SuppressCompare:           settings.SuppressCompareAssertion,
			SuppressAsync:             settings.SuppressAsyncAssertion,
			ForbidFocus:               settings.ForbidFocusContainer,
			SuppressTypeCompare:       settings.SuppressTypeCompareWarning,
			AllowHaveLen0:             settings.AllowHaveLenZero,
			ForceExpectTo:             settings.ForceExpectTo,
			ValidateAsyncIntervals:    settings.ValidateAsyncIntervals,
			ForbidSpecPollution:       settings.ForbidSpecPollution,
			ForceSucceedForFuncs:      settings.ForceSucceedForFuncs,
			ForceAssertionDescription: settings.ForceAssertionDescription,
			ForeToNot:                 settings.ForeToNot,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(ginkgolinter.NewAnalyzerWithConfig(cfg)).
		WithDesc("enforces standards of using ginkgo and gomega").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

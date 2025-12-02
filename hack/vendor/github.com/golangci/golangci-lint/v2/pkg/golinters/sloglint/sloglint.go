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

package sloglint

import (
	"go-simpler.org/sloglint"

	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.SlogLintSettings) *goanalysis.Linter {
	var opts *sloglint.Options

	if settings != nil {
		opts = &sloglint.Options{
			NoMixedArgs:    settings.NoMixedArgs,
			KVOnly:         settings.KVOnly,
			AttrOnly:       settings.AttrOnly,
			NoGlobal:       settings.NoGlobal,
			ContextOnly:    settings.Context,
			StaticMsg:      settings.StaticMsg,
			MsgStyle:       settings.MsgStyle,
			NoRawKeys:      settings.NoRawKeys,
			KeyNamingCase:  settings.KeyNamingCase,
			ForbiddenKeys:  settings.ForbiddenKeys,
			ArgsOnSepLines: settings.ArgsOnSepLines,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(sloglint.New(opts)).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}

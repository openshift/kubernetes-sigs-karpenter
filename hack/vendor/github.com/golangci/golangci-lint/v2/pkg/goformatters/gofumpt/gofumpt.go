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

package gofumpt

import (
	"strings"

	gofumpt "mvdan.cc/gofumpt/format"

	"github.com/golangci/golangci-lint/v2/pkg/config"
)

const Name = "gofumpt"

type Formatter struct {
	options gofumpt.Options
}

func New(settings *config.GoFumptSettings, goVersion string) *Formatter {
	var options gofumpt.Options

	if settings != nil {
		options = gofumpt.Options{
			LangVersion: getLangVersion(goVersion),
			ModulePath:  settings.ModulePath,
			ExtraRules:  settings.ExtraRules,
		}
	}

	return &Formatter{options: options}
}

func (*Formatter) Name() string {
	return Name
}

func (f *Formatter) Format(_ string, src []byte) ([]byte, error) {
	return gofumpt.Source(src, f.options)
}

func getLangVersion(v string) string {
	return "go" + strings.TrimPrefix(v, "go")
}

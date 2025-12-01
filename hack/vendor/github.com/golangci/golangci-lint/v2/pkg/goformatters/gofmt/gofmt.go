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

package gofmt

import (
	"github.com/golangci/gofmt/gofmt"

	"github.com/golangci/golangci-lint/v2/pkg/config"
)

const Name = "gofmt"

type Formatter struct {
	options gofmt.Options
}

func New(settings *config.GoFmtSettings) *Formatter {
	options := gofmt.Options{}

	if settings != nil {
		options.NeedSimplify = settings.Simplify

		for _, rule := range settings.RewriteRules {
			options.RewriteRules = append(options.RewriteRules, gofmt.RewriteRule(rule))
		}
	}

	return &Formatter{options: options}
}

func (*Formatter) Name() string {
	return Name
}

func (f *Formatter) Format(filename string, src []byte) ([]byte, error) {
	return gofmt.Source(filename, src, f.options)
}

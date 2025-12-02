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

package goimports

import (
	"strings"

	"golang.org/x/tools/imports"

	"github.com/golangci/golangci-lint/v2/pkg/config"
)

const Name = "goimports"

type Formatter struct{}

func New(settings *config.GoImportsSettings) *Formatter {
	if settings != nil {
		imports.LocalPrefix = strings.Join(settings.LocalPrefixes, ",")
	}

	return &Formatter{}
}

func (*Formatter) Name() string {
	return Name
}

func (*Formatter) Format(filename string, src []byte) ([]byte, error) {
	// The `imports.LocalPrefix` (`settings.LocalPrefixes`) is a global var.
	return imports.Process(filename, src, nil)
}

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

package section

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/mod/modfile"

	"github.com/daixiang0/gci/pkg/parse"
	"github.com/daixiang0/gci/pkg/specificity"
)

const LocalModuleType = "localmodule"

type LocalModule struct {
	Path string
}

func (m *LocalModule) MatchSpecificity(spec *parse.GciImports) specificity.MatchSpecificity {
	if spec.Path == m.Path || strings.HasPrefix(spec.Path, m.Path+"/") {
		return specificity.LocalModule{}
	}

	return specificity.MisMatch{}
}

func (m *LocalModule) String() string {
	return LocalModuleType
}

func (m *LocalModule) Type() string {
	return LocalModuleType
}

// Configure configures the module section by finding the module
// for the current path
func (m *LocalModule) Configure(path string) error {
	if path != "" {
		m.Path = path
	} else {
		path, err := findLocalModule()
		if err != nil {
			return fmt.Errorf("finding local modules for `localModule` configuration: %w", err)
		}
		m.Path = path
	}

	return nil
}

func findLocalModule() (string, error) {
	b, err := os.ReadFile("go.mod")
	if err != nil {
		return "", fmt.Errorf("reading go.mod: %w", err)
	}

	return modfile.ModulePath(b), nil
}

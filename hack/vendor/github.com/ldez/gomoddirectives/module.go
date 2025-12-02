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

package gomoddirectives

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ldez/grignotin/goenv"
	"golang.org/x/mod/modfile"
)

// GetModuleFile gets module file.
func GetModuleFile() (*modfile.File, error) {
	goMod, err := goenv.GetOne(context.Background(), goenv.GOMOD)
	if err != nil {
		return nil, err
	}

	mod, err := parseGoMod(goMod)
	if err != nil {
		return nil, fmt.Errorf("failed to parse go.mod (%s): %w", goMod, err)
	}

	return mod, nil
}

func parseGoMod(goMod string) (*modfile.File, error) {
	raw, err := os.ReadFile(filepath.Clean(goMod))
	if err != nil {
		return nil, fmt.Errorf("reading go.mod file: %w", err)
	}

	return modfile.Parse("go.mod", raw, nil)
}

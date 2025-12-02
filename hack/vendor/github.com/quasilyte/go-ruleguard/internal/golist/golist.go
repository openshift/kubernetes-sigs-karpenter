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

package golist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
)

// Package is `go list --json` output structure.
type Package struct {
	Dir        string   // directory containing package sources
	ImportPath string   // import path of package in dir
	GoFiles    []string // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
}

// JSON runs `go list --json` for the specified pkgName and returns the parsed JSON.
func JSON(pkgPath string) (*Package, error) {
	out, err := exec.Command("go", "list", "--json", pkgPath).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("go list error (%v): %s", err, out)
	}

	var pkg Package
	if err := json.NewDecoder(bytes.NewReader(out)).Decode(&pkg); err != io.EOF && err != nil {
		return nil, err
	}
	return &pkg, nil
}

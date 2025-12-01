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

package goenv

import (
	"errors"
	"os/exec"
	"strings"
)

func Read() (map[string]string, error) {
	// pass in a fixed set of var names to avoid needing to unescape output
	// pass in literals here instead of a variable list to avoid security linter warnings about command injection
	out, err := exec.Command("go", "env", "GOROOT", "GOPATH", "GOARCH", "GOOS", "CGO_ENABLED").CombinedOutput()
	if err != nil {
		return nil, err
	}
	return parseGoEnv([]string{"GOROOT", "GOPATH", "GOARCH", "GOOS", "CGO_ENABLED"}, out)
}

func parseGoEnv(varNames []string, data []byte) (map[string]string, error) {
	vars := make(map[string]string)

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	for i, varName := range varNames {
		if i < len(lines) && len(lines[i]) > 0 {
			vars[varName] = lines[i]
		}
	}

	if len(vars) == 0 {
		return nil, errors.New("empty env set")
	}

	return vars, nil
}

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

package musttag

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func getMainModule() (string, error) {
	args := [...]string{"go", "mod", "edit", "-json"}

	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return "", fmt.Errorf("running %q: %w", strings.Join(args[:], " "), err)
	}

	var info struct {
		Module struct {
			Path string `json:"Path"`
		} `json:"Module"`
	}
	if err := json.Unmarshal(out, &info); err != nil {
		return "", fmt.Errorf("decoding module info: %w\n%s", err, out)
	}

	return info.Module.Path, nil
}

// based on golang.org/x/tools/imports.VendorlessPath
func cutVendor(path string) string {
	var prefix string

	switch {
	case strings.HasPrefix(path, "(*"):
		prefix, path = "(*", path[len("(*"):]
	case strings.HasPrefix(path, "("):
		prefix, path = "(", path[len("("):]
	}

	if i := strings.LastIndex(path, "/vendor/"); i >= 0 {
		return prefix + path[i+len("/vendor/"):]
	}
	if strings.HasPrefix(path, "vendor/") {
		return prefix + path[len("vendor/"):]
	}

	return prefix + path
}

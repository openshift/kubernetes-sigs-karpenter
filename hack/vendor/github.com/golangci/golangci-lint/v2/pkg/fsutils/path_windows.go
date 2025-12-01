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

//go:build windows

package fsutils

import (
	"path/filepath"
	"regexp"
	"strings"
)

var separatorToReplace = regexp.QuoteMeta(string(filepath.Separator))

// NormalizePathInRegex normalizes path in regular expressions.
// noop on Unix.
// This replacing should be safe because "/" are disallowed in Windows
// https://docs.microsoft.com/windows/win32/fileio/naming-a-file
func NormalizePathInRegex(path string) string {
	// remove redundant character escape "\/" https://github.com/golangci/golangci-lint/issues/3277
	clean := regexp.MustCompile(`\\+/`).
		ReplaceAllStringFunc(path, func(s string) string {
			if strings.Count(s, "\\")%2 == 0 {
				return s
			}
			return s[1:]
		})

	return strings.ReplaceAll(clean, "/", separatorToReplace)
}

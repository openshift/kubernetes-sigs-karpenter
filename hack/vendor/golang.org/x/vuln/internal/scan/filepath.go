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

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scan

import (
	"path/filepath"
	"strings"
)

// AbsRelShorter takes path and returns its path relative
// to the current directory, if shorter. Returns path
// when path is an empty string or upon any error.
func AbsRelShorter(path string) string {
	if path == "" {
		return ""
	}

	c, err := filepath.Abs(".")
	if err != nil {
		return path
	}
	r, err := filepath.Rel(c, path)
	if err != nil {
		return path
	}

	rSegments := strings.Split(r, string(filepath.Separator))
	pathSegments := strings.Split(path, string(filepath.Separator))
	if len(rSegments) < len(pathSegments) {
		return r
	}
	return path
}

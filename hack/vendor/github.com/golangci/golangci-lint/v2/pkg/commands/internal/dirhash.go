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

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/sumdb/dirhash"
)

// Slightly modified copy of [dirhash.HashDir].
// https://github.com/golang/mod/blob/v0.28.0/sumdb/dirhash/hash.go#L67-L79
func hashDir(dir, prefix string, hash dirhash.Hash) (string, error) {
	files, err := dirFiles(dir, prefix)
	if err != nil {
		return "", err
	}

	osOpen := func(name string) (io.ReadCloser, error) {
		return os.Open(filepath.Join(dir, strings.TrimPrefix(name, prefix)))
	}

	return hash(files, osOpen)
}

// Modified copy of [dirhash.DirFiles].
// https://github.com/golang/mod/blob/v0.28.0/sumdb/dirhash/hash.go#L81-L109
// And adapted to globally follows the rules from https://github.com/golang/mod/blob/v0.28.0/zip/zip.go
func dirFiles(dir, prefix string) ([]string, error) {
	var files []string

	dir = filepath.Clean(dir)

	err := filepath.Walk(dir, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if dir == file {
				// Don't skip the top-level directory.
				return nil
			}

			switch info.Name() {
			// Skip vendor and node directories.
			case "vendor", "node_modules":
				return filepath.SkipDir

			// Skip VCS directories.
			case ".bzr", ".git", ".hg", ".svn":
				return filepath.SkipDir
			}

			// Skip submodules (directories containing go.mod files).
			if goModInfo, err := os.Lstat(filepath.Join(dir, "go.mod")); err == nil && !goModInfo.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		if file == dir {
			return fmt.Errorf("%s is not a directory", dir)
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		rel := file

		if dir != "." {
			rel = file[len(dir)+1:]
		}

		f := filepath.Join(prefix, rel)

		files = append(files, filepath.ToSlash(f))

		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

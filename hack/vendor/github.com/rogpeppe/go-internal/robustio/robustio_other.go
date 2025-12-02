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

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !darwin

package robustio

import (
	"os"
)

func rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func readFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func removeAll(path string) error {
	return os.RemoveAll(path)
}

func isEphemeralError(err error) bool {
	return false
}

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

package io

import (
	"io/ioutil"
	"os"
)

type stdInFile struct{}

func (s stdInFile) Load() ([]byte, error) {
	return ioutil.ReadAll(os.Stdin)
}

func (s stdInFile) Path() string {
	return "StdIn"
}

var StdInGenerator FileGeneratorFunc = func() ([]FileObj, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return []FileObj{stdInFile{}}, nil
	}
	return []FileObj{}, nil
}

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

package in_toto

import (
	"errors"
	"os"
)

func isWritable(path string) error {
	// get fileInfo
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	// check if path is a directory
	if !info.IsDir() {
		return errors.New("not a directory")
	}

	// Check if the user bit is enabled in file permission
	if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		return errors.New("not writable")
	}
	return nil
}

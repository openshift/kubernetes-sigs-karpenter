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

package ini

import "fmt"

// UnableToReadFile is an error indicating that a ini file could not be read
type UnableToReadFile struct {
	Err error
}

// Error returns an error message and the underlying error message if present
func (e *UnableToReadFile) Error() string {
	base := "unable to read file"
	if e.Err == nil {
		return base
	}
	return fmt.Sprintf("%s: %v", base, e.Err)
}

// Unwrap returns the underlying error
func (e *UnableToReadFile) Unwrap() error {
	return e.Err
}

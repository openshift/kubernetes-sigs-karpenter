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

package retry

import "fmt"

// MaxAttemptsError provides the error when the maximum number of attempts have
// been exceeded.
type MaxAttemptsError struct {
	Attempt int
	Err     error
}

func (e *MaxAttemptsError) Error() string {
	return fmt.Sprintf("exceeded maximum number of attempts, %d, %v", e.Attempt, e.Err)
}

// Unwrap returns the nested error causing the max attempts error. Provides the
// implementation for errors.Is and errors.As to unwrap nested errors.
func (e *MaxAttemptsError) Unwrap() error {
	return e.Err
}

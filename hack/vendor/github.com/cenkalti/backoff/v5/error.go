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

package backoff

import (
	"fmt"
	"time"
)

// PermanentError signals that the operation should not be retried.
type PermanentError struct {
	Err error
}

// Permanent wraps the given err in a *PermanentError.
func Permanent(err error) error {
	if err == nil {
		return nil
	}
	return &PermanentError{
		Err: err,
	}
}

// Error returns a string representation of the Permanent error.
func (e *PermanentError) Error() string {
	return e.Err.Error()
}

// Unwrap returns the wrapped error.
func (e *PermanentError) Unwrap() error {
	return e.Err
}

// RetryAfterError signals that the operation should be retried after the given duration.
type RetryAfterError struct {
	Duration time.Duration
}

// RetryAfter returns a RetryAfter error that specifies how long to wait before retrying.
func RetryAfter(seconds int) error {
	return &RetryAfterError{Duration: time.Duration(seconds) * time.Second}
}

// Error returns a string representation of the RetryAfter error.
func (e *RetryAfterError) Error() string {
	return fmt.Sprintf("retry after %s", e.Duration)
}

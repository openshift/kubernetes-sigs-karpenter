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

package errs

import (
	"fmt"
	"io"
)

// Group is a list of errors.
type Group []error

// Combine combines multiple non-empty errors into a single error.
func Combine(errs ...error) error {
	var group Group
	group.Add(errs...)
	return group.Err()
}

// Add adds non-empty errors to the Group.
func (group *Group) Add(errs ...error) {
	for _, err := range errs {
		if err != nil {
			*group = append(*group, err)
		}
	}
}

// Err returns an error containing all of the non-nil errors.
// If there was only one error, it will return it.
// If there were none, it returns nil.
func (group Group) Err() error {
	sanitized := group.sanitize()
	if len(sanitized) == 0 {
		return nil
	}
	if len(sanitized) == 1 {
		return sanitized[0]
	}
	return combinedError(sanitized)
}

// sanitize returns group that doesn't contain nil-s
func (group Group) sanitize() Group {
	// sanity check for non-nil errors
	for i, err := range group {
		if err == nil {
			sanitized := make(Group, 0, len(group)-1)
			sanitized = append(sanitized, group[:i]...)
			sanitized.Add(group[i+1:]...)
			return sanitized
		}
	}

	return group
}

// combinedError is a list of non-empty errors
type combinedError []error

// Unwrap returns the first error.
func (group combinedError) Unwrap() []error { return group }

// Error returns error string delimited by semicolons.
func (group combinedError) Error() string { return fmt.Sprintf("%v", group) }

// Format handles the formatting of the error. Using a "+" on the format
// string specifier will cause the errors to be formatted with "+" and
// delimited by newlines. They are delimited by semicolons otherwise.
func (group combinedError) Format(f fmt.State, c rune) {
	delim := "; "
	if f.Flag(int('+')) {
		io.WriteString(f, "group:\n--- ")
		delim = "\n--- "
	}

	for i, err := range group {
		if i != 0 {
			io.WriteString(f, delim)
		}
		if formatter, ok := err.(fmt.Formatter); ok {
			formatter.Format(f, c)
		} else {
			fmt.Fprintf(f, "%v", err)
		}
	}
}

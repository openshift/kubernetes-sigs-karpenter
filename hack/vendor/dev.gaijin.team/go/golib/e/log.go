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

package e

import (
	"dev.gaijin.team/go/golib/fields"
)

// ErrorLogger defines a function that logs an error message, an error, and optional fields.
type ErrorLogger func(msg string, err error, fs ...fields.Field)

// Log logs the provided error using the given ErrorLogger function.
//
// If err is nil, Log does nothing. If err is of type Err, its reason is used as the log message,
// the wrapped error is passed as the error, and its fields are passed as log fields.
// For other error types, err.Error() is used as the message and nil is passed as the error.
func Log(err error, f ErrorLogger) {
	if err == nil {
		return
	}

	// We're not interested in wrapped error, therefore we're only typecasting it.
	if e, ok := err.(*Err); ok { //nolint:errorlint
		var wrapped error
		if len(e.errs) > 1 {
			wrapped = e.errs[1]
		}

		f(e.Reason(), wrapped, e.fields...)

		return
	}

	f(err.Error(), nil)
}

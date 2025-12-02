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

package chroma

import (
	"io"
)

// A Formatter for Chroma lexers.
type Formatter interface {
	// Format returns a formatting function for tokens.
	//
	// If the iterator panics, the Formatter should recover.
	Format(w io.Writer, style *Style, iterator Iterator) error
}

// A FormatterFunc is a Formatter implemented as a function.
//
// Guards against iterator panics.
type FormatterFunc func(w io.Writer, style *Style, iterator Iterator) error

func (f FormatterFunc) Format(w io.Writer, s *Style, it Iterator) (err error) { // nolint
	defer func() {
		if perr := recover(); perr != nil {
			err = perr.(error)
		}
	}()
	return f(w, s, it)
}

type recoveringFormatter struct {
	Formatter
}

func (r recoveringFormatter) Format(w io.Writer, s *Style, it Iterator) (err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = perr.(error)
		}
	}()
	return r.Formatter.Format(w, s, it)
}

// RecoveringFormatter wraps a formatter with panic recovery.
func RecoveringFormatter(formatter Formatter) Formatter { return recoveringFormatter{formatter} }

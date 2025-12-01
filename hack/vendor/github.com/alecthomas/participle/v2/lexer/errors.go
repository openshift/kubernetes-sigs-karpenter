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

package lexer

import "fmt"

// This file exists to break circular imports. The types and functions in here
// mirror those in the participle package.

type errorInterface interface {
	error
	Message() string
	Position() Position
}

// Error represents an error while lexing.
//
// It complies with the participle.Error interface.
type Error struct {
	Msg string
	Pos Position
}

var _ errorInterface = &Error{}

// Creates a new Error at the given position.
func errorf(pos Position, format string, args ...interface{}) *Error {
	return &Error{Msg: fmt.Sprintf(format, args...), Pos: pos}
}

func (e *Error) Message() string    { return e.Msg } // nolint: golint
func (e *Error) Position() Position { return e.Pos } // nolint: golint

// Error formats the error with FormatError.
func (e *Error) Error() string { return formatError(e.Pos, e.Msg) }

// An error in the form "[<filename>:][<line>:<pos>:] <message>"
func formatError(pos Position, message string) string {
	msg := ""
	if pos.Filename != "" {
		msg += pos.Filename + ":"
	}
	if pos.Line != 0 || pos.Column != 0 {
		msg += fmt.Sprintf("%d:%d:", pos.Line, pos.Column)
	}
	if msg != "" {
		msg += " " + message
	} else {
		msg = message
	}
	return msg
}

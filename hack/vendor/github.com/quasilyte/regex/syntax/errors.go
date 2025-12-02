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

package syntax

type ParseError struct {
	Pos     Position
	Message string
}

func (e ParseError) Error() string { return e.Message }

func throw(pos Position, message string) {
	panic(ParseError{Pos: pos, Message: message})
}

func throwExpectedFound(pos Position, expected, found string) {
	throw(pos, "expected '"+expected+"', found '"+found+"'")
}

func throwUnexpectedToken(pos Position, token string) {
	throw(pos, "unexpected token: "+token)
}

func newPos(begin, end int) Position {
	return Position{
		Begin: uint16(begin),
		End:   uint16(end),
	}
}

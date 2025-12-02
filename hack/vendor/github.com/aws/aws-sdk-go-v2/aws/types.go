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

package aws

import (
	"fmt"
)

// Ternary is an enum allowing an unknown or none state in addition to a bool's
// true and false.
type Ternary int

func (t Ternary) String() string {
	switch t {
	case UnknownTernary:
		return "unknown"
	case FalseTernary:
		return "false"
	case TrueTernary:
		return "true"
	default:
		return fmt.Sprintf("unknown value, %d", int(t))
	}
}

// Bool returns true if the value is TrueTernary, false otherwise.
func (t Ternary) Bool() bool {
	return t == TrueTernary
}

// Enumerations for the values of the Ternary type.
const (
	UnknownTernary Ternary = iota
	FalseTernary
	TrueTernary
)

// BoolTernary returns a true or false Ternary value for the bool provided.
func BoolTernary(v bool) Ternary {
	if v {
		return TrueTernary
	}
	return FalseTernary
}

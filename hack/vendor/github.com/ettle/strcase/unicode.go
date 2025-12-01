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

package strcase

import "unicode"

// Unicode functions, optimized for the common case of ascii
// No performance lost by wrapping since these functions get inlined by the compiler

func isUpper(r rune) bool {
	return unicode.IsUpper(r)
}

func isLower(r rune) bool {
	return unicode.IsLower(r)
}

func isNumber(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return unicode.IsNumber(r)
}

func isSpace(r rune) bool {
	if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
		return true
	} else if r < 128 {
		return false
	}
	return unicode.IsSpace(r)
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	} else if r < 128 {
		return r
	}
	return unicode.ToUpper(r)
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	} else if r < 128 {
		return r
	}
	return unicode.ToLower(r)
}

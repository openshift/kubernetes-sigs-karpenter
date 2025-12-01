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

package textmatch

import "regexp"

// Pattern is a compiled regular expression.
type Pattern interface {
	MatchString(s string) bool
	Match(b []byte) bool
}

// Compile parses a regular expression and returns a compiled
// pattern that can match inputs described by the regexp.
//
// Semantically it's close to the regexp.Compile, but
// it does recognize some common patterns and creates
// a more optimized matcher for them.
func Compile(re string) (Pattern, error) {
	return compile(re)
}

// IsRegexp reports whether p is implemented using regexp.
// False means that the underlying matcher is something optimized.
func IsRegexp(p Pattern) bool {
	_, ok := p.(*regexp.Regexp)
	return ok
}

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

package exhaustive

import (
	"flag"
	"regexp"
	"strings"
)

var _ flag.Value = (*regexpFlag)(nil)
var _ flag.Value = (*stringsFlag)(nil)

// regexpFlag implements flag.Value for parsing
// regular expression flag inputs.
type regexpFlag struct{ re *regexp.Regexp }

func (f *regexpFlag) String() string {
	if f == nil || f.re == nil {
		return ""
	}
	return f.re.String()
}

func (f *regexpFlag) Set(expr string) error {
	if expr == "" {
		f.re = nil
		return nil
	}

	re, err := regexp.Compile(expr)
	if err != nil {
		return err
	}

	f.re = re
	return nil
}

// stringsFlag implements flag.Value for parsing a comma-separated string
// list. Surrounding whitespace is stripped from the input and from each
// element. If filter is non-nil it is called for each element in the input.
type stringsFlag struct {
	elements []string
	filter   func(string) error
}

func (f *stringsFlag) String() string {
	if f == nil {
		return ""
	}
	return strings.Join(f.elements, ",")
}

func (f *stringsFlag) filterFunc() func(string) error {
	if f.filter != nil {
		return f.filter
	}
	return func(_ string) error { return nil }
}

func (f *stringsFlag) Set(input string) error {
	input = strings.TrimSpace(input)
	if input == "" {
		f.elements = nil
		return nil
	}

	for _, el := range strings.Split(input, ",") {
		el = strings.TrimSpace(el)
		if err := f.filterFunc()(el); err != nil {
			return err
		}
		f.elements = append(f.elements, el)
	}
	return nil
}

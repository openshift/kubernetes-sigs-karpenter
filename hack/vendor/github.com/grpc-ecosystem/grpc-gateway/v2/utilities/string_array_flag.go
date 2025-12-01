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

package utilities

import (
	"flag"
	"strings"
)

// flagInterface is a cut down interface to `flag`
type flagInterface interface {
	Var(value flag.Value, name string, usage string)
}

// StringArrayFlag defines a flag with the specified name and usage string.
// The return value is the address of a `StringArrayFlags` variable that stores the repeated values of the flag.
func StringArrayFlag(f flagInterface, name string, usage string) *StringArrayFlags {
	value := &StringArrayFlags{}
	f.Var(value, name, usage)
	return value
}

// StringArrayFlags is a wrapper of `[]string` to provider an interface for `flag.Var`
type StringArrayFlags []string

// String returns a string representation of `StringArrayFlags`
func (i *StringArrayFlags) String() string {
	return strings.Join(*i, ",")
}

// Set appends a value to `StringArrayFlags`
func (i *StringArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

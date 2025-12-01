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

package config

import (
	"regexp"
	"sync"

	"github.com/butuzov/ireturn/analyzer/internal/types"
)

// defaultConfig is core of the validation, ...
// todo(butuzov): write proper intro...

type defaultConfig struct {
	List []string

	// private fields (for search optimization look ups)
	once  sync.Once
	quick uint8
	list  []*regexp.Regexp
}

func (config *defaultConfig) Has(i types.IFace) bool {
	config.once.Do(config.compileList)

	if config.quick&uint8(i.Type) > 0 {
		return true
	}

	// not a named interface (because error, interface{}, anon interface has keywords.)
	if i.Type&types.NamedInterface == 0 && i.Type&types.NamedStdInterface == 0 {
		return false
	}

	for _, re := range config.list {
		if re.MatchString(i.Name) {
			return true
		}
	}

	return false
}

// compileList will transform text list into a bitmask for quick searches and
// slice of regular expressions for quick searches.
func (config *defaultConfig) compileList() {
	for _, str := range config.List {
		switch str {
		case types.NameError:
			config.quick |= uint8(types.ErrorInterface)
		case types.NameEmpty:
			config.quick |= uint8(types.EmptyInterface)
		case types.NameAnon:
			config.quick |= uint8(types.AnonInterface)
		case types.NameStdLib:
			config.quick |= uint8(types.NamedStdInterface)
		case types.NameGeneric:
			config.quick |= uint8(types.Generic)
		}

		// allow to parse regular expressions
		// todo(butuzov): how can we log error in golangci-lint?
		if re, err := regexp.Compile(str); err == nil {
			config.list = append(config.list, re)
		}
	}
}

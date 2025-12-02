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

package analyzer

import (
	"fmt"
	"sort"
	"strings"
)

func newDefaultCheckedTypes() checkedTypes {
	return checkedTypes{
		chanType:      {},
		funcType:      {},
		ifaceType:     {},
		mapType:       {},
		ptrType:       {},
		uintptrType:   {},
		unsafeptrType: {},
	}
}

const separator = ','

type typeName string

func (t typeName) S() string {
	return string(t)
}

const (
	ptrType       typeName = "ptr"
	funcType      typeName = "func"
	ifaceType     typeName = "iface"
	mapType       typeName = "map"
	chanType      typeName = "chan"
	uintptrType   typeName = "uintptr"
	unsafeptrType typeName = "unsafeptr"
)

type checkedTypes map[typeName]struct{}

func (c checkedTypes) Contains(t typeName) bool {
	_, ok := c[t]
	return ok
}

func (c checkedTypes) String() string {
	result := make([]string, 0, len(c))
	for t := range c {
		result = append(result, t.S())
	}

	sort.Strings(result)
	return strings.Join(result, string(separator))
}

func (c checkedTypes) Set(s string) error {
	types := strings.FieldsFunc(s, func(c rune) bool { return c == separator })
	if len(types) == 0 {
		return nil
	}

	c.disableAll()
	for _, t := range types {
		switch tt := typeName(t); tt {
		case ptrType, funcType, ifaceType, mapType, chanType, uintptrType, unsafeptrType:
			c[tt] = struct{}{}
		default:
			return fmt.Errorf("unknown checked type name %q (see help)", t)
		}
	}

	return nil
}

func (c checkedTypes) disableAll() {
	for k := range c {
		delete(c, k)
	}
}

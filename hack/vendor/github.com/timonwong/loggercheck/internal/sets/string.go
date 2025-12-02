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

package sets

import (
	"sort"
	"strings"
)

type Empty struct{}

type StringSet map[string]Empty

func NewString(items ...string) StringSet {
	s := make(StringSet)
	s.Insert(items...)
	return s
}

func (s StringSet) Insert(items ...string) {
	for _, item := range items {
		s[item] = Empty{}
	}
}

func (s StringSet) Has(item string) bool {
	_, contained := s[item]
	return contained
}

func (s StringSet) List() []string {
	if len(s) == 0 {
		return nil
	}

	res := make([]string, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	sort.Strings(res)
	return res
}

// Set implements flag.Value interface.
func (s *StringSet) Set(v string) error {
	v = strings.TrimSpace(v)
	if v == "" {
		*s = nil
		return nil
	}

	parts := strings.Split(v, ",")
	set := NewString(parts...)
	*s = set
	return nil
}

// String implements flag.Value interface
func (s StringSet) String() string {
	return strings.Join(s.List(), ",")
}

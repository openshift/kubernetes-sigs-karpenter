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

package parse

import (
	"strings"
)

type Env []string

func (e Env) Get(name string) string {
	v, _ := e.Lookup(name)
	return v
}

func (e Env) Has(name string) bool {
	_, ok := e.Lookup(name)
	return ok
}

func (e Env) Lookup(name string) (string, bool) {
	prefix := name + "="
	for _, pair := range e {
		if strings.HasPrefix(pair, prefix) {
			return pair[len(prefix):], true
		}
	}
	return "", false
}

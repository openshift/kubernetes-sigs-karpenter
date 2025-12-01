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

package yit

import "gopkg.in/yaml.v3"

func (next Iterator) AnyMatch(p Predicate) bool {
	iterator := next.Filter(p)
	_, ok := iterator()
	return ok
}

func (next Iterator) AllMatch(p Predicate) bool {
	result := true
	for node, ok := next(); ok && result; node, ok = next() {
		result = result && p(node)
	}

	return result
}

func (next Iterator) ToArray() (result []*yaml.Node) {
	for node, ok := next(); ok; node, ok = next() {
		result = append(result, node)
	}
	return
}

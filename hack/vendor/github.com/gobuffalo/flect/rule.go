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

package flect

type ruleFn func(string) string

type rule struct {
	suffix string
	fn     ruleFn
}

func simpleRuleFunc(suffix, repl string) func(string) string {
	return func(s string) string {
		s = s[:len(s)-len(suffix)]
		return s + repl
	}
}

func noop(s string) string { return s }

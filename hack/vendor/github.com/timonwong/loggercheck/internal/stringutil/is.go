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

package stringutil

import "unicode/utf8"

// IsASCII returns true if string are ASCII.
func IsASCII(s string) bool {
	for _, r := range s {
		if r >= utf8.RuneSelf {
			// Not ASCII.
			return false
		}
	}

	return true
}

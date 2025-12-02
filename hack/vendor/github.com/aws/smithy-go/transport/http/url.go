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

package http

import "strings"

// JoinPath returns an absolute URL path composed of the two paths provided.
// Enforces that the returned path begins with '/'. If added path is empty the
// returned path suffix will match the first parameter suffix.
func JoinPath(a, b string) string {
	if len(a) == 0 {
		a = "/"
	} else if a[0] != '/' {
		a = "/" + a
	}

	if len(b) != 0 && b[0] == '/' {
		b = b[1:]
	}

	if len(b) != 0 && len(a) > 1 && a[len(a)-1] != '/' {
		a = a + "/"
	}

	return a + b
}

// JoinRawQuery returns an absolute raw query expression. Any duplicate '&'
// will be collapsed to single separator between values.
func JoinRawQuery(a, b string) string {
	a = strings.TrimFunc(a, isAmpersand)
	b = strings.TrimFunc(b, isAmpersand)

	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	return a + "&" + b
}

func isAmpersand(v rune) bool {
	return v == '&'
}

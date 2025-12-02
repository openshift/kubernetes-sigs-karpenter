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

package locafero

import "fmt"

// NameWithExtensions creates a list of names from a base name and a list of extensions.
//
// TODO: find a better name for this function.
func NameWithExtensions(baseName string, extensions ...string) []string {
	var names []string

	if baseName == "" {
		return names
	}

	for _, ext := range extensions {
		if ext == "" {
			continue
		}

		names = append(names, fmt.Sprintf("%s.%s", baseName, ext))
	}

	return names
}

// NameWithOptionalExtensions creates a list of names from a base name and a list of extensions,
// plus it adds the base name (without any extensions) to the end of the list.
//
// TODO: find a better name for this function.
func NameWithOptionalExtensions(baseName string, extensions ...string) []string {
	var names []string

	if baseName == "" {
		return names
	}

	names = NameWithExtensions(baseName, extensions...)
	names = append(names, baseName)

	return names
}

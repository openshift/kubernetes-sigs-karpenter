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

package roles

import (
	"strconv"
	"strings"
)

var TopLevelRoles = map[string]struct{}{
	"root":      {},
	"targets":   {},
	"snapshot":  {},
	"timestamp": {},
}

func IsTopLevelRole(name string) bool {
	_, ok := TopLevelRoles[name]
	return ok
}

func IsDelegatedTargetsRole(name string) bool {
	return !IsTopLevelRole(name)
}

func IsTopLevelManifest(name string) bool {
	if IsVersionedManifest(name) {
		var found bool
		_, name, found = strings.Cut(name, ".")
		if !found {
			panic("expected a versioned manifest of the form x.role.json")
		}
	}
	return IsTopLevelRole(strings.TrimSuffix(name, ".json"))
}

func IsDelegatedTargetsManifest(name string) bool {
	return !IsTopLevelManifest(name)
}

func IsVersionedManifest(name string) bool {
	parts := strings.Split(name, ".")
	// Versioned manifests have the form "x.role.json"
	if len(parts) < 3 {
		return false
	}

	_, err := strconv.Atoi(parts[0])
	return err == nil
}

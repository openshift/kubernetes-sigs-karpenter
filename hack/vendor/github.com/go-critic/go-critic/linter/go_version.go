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

package linter

import (
	"fmt"
	"strconv"
	"strings"
)

type GoVersion struct {
	Major int
	Minor int
}

// GreaterOrEqual performs $v >= $other operation.
//
// In other words, it reports whether $v version constraint can use
// a feature from the $other Go version.
//
// As a special case, Major=0 covers all versions.
func (v GoVersion) GreaterOrEqual(other GoVersion) bool {
	switch {
	case v.Major == 0:
		return true
	case v.Major == other.Major:
		return v.Minor >= other.Minor
	default:
		return v.Major >= other.Major
	}
}

func ParseGoVersion(version string) (GoVersion, error) {
	var result GoVersion
	version = strings.TrimPrefix(version, "go")
	if version == "" {
		return result, nil
	}
	parts := strings.Split(version, ".")
	if len(parts) != 2 {
		return result, fmt.Errorf("invalid Go version format: %s", version)
	}
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return result, fmt.Errorf("invalid major version part: %s: %w", parts[0], err)
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return result, fmt.Errorf("invalid minor version part: %s: %w", parts[1], err)
	}
	result.Major = major
	result.Minor = minor
	return result, nil
}

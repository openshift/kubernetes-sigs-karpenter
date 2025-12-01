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

package filters

import (
	"encoding/json"

	"github.com/docker/docker/api/types/versions"
)

// ToParamWithVersion encodes Args as a JSON string. If version is less than 1.22
// then the encoded format will use an older legacy format where the values are a
// list of strings, instead of a set.
//
// Deprecated: do not use in any new code; use ToJSON instead
func ToParamWithVersion(version string, a Args) (string, error) {
	out, err := ToJSON(a)
	if out == "" || err != nil {
		return "", nil
	}
	if version != "" && versions.LessThan(version, "1.22") {
		return encodeLegacyFilters(out)
	}
	return out, nil
}

// encodeLegacyFilters encodes Args in the legacy format as used in API v1.21 and older.
// where values are a list of strings, instead of a set.
//
// Don't use in any new code; use [filters.ToJSON]] instead.
func encodeLegacyFilters(currentFormat string) (string, error) {
	// The Args.fields field is not exported, but used to marshal JSON,
	// so we'll marshal to the new format, then unmarshal to get the
	// fields, and marshal again.
	//
	// This is far from optimal, but this code is only used for deprecated
	// API versions, so should not be hit commonly.
	var argsFields map[string]map[string]bool
	err := json.Unmarshal([]byte(currentFormat), &argsFields)
	if err != nil {
		return "", err
	}

	buf, err := json.Marshal(convertArgsToSlice(argsFields))
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func convertArgsToSlice(f map[string]map[string]bool) map[string][]string {
	m := map[string][]string{}
	for k, v := range f {
		values := []string{}
		for kk := range v {
			if v[kk] {
				values = append(values, kk)
			}
		}
		m[k] = values
	}
	return m
}

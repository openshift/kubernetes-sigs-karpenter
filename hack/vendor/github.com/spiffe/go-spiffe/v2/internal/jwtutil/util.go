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

package jwtutil

import (
	"crypto"

	"github.com/spiffe/go-spiffe/v2/internal/cryptoutil"
)

// CopyJWTAuthorities copies JWT authorities from a map to a new map.
func CopyJWTAuthorities(jwtAuthorities map[string]crypto.PublicKey) map[string]crypto.PublicKey {
	copiedJWTAuthorities := make(map[string]crypto.PublicKey)
	for key, jwtAuthority := range jwtAuthorities {
		copiedJWTAuthorities[key] = jwtAuthority
	}
	return copiedJWTAuthorities
}

func JWTAuthoritiesEqual(a, b map[string]crypto.PublicKey) bool {
	if len(a) != len(b) {
		return false
	}

	for k, pka := range a {
		pkb, ok := b[k]
		if !ok {
			return false
		}
		if equal, _ := cryptoutil.PublicKeyEqual(pka, pkb); !equal {
			return false
		}
	}

	return true
}

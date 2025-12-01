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

package in_toto

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

/*
getHashMapping returns a mapping from hash algorithm to supported hash
interface.
*/
func getHashMapping() map[string]func() hash.Hash {
	return map[string]func() hash.Hash{
		"sha256": sha256.New,
		"sha512": sha512.New,
		"sha384": sha512.New384,
	}
}

/*
hashToHex calculates the hash over data based on hash algorithm h.
*/
func hashToHex(h hash.Hash, data []byte) []byte {
	h.Write(data)
	// We need to use h.Sum(nil) here, because otherwise hash.Sum() appends
	// the hash to the passed data. So instead of having only the hash
	// we would get: "dataHASH"
	return h.Sum(nil)
}

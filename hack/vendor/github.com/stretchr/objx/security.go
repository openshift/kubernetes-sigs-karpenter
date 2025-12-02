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

package objx

import (
	"crypto/sha1"
	"encoding/hex"
)

// HashWithKey hashes the specified string using the security key
func HashWithKey(data, key string) string {
	d := sha1.Sum([]byte(data + ":" + key))
	return hex.EncodeToString(d[:])
}

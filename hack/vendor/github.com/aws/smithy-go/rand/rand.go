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

package rand

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

func init() {
	Reader = rand.Reader
}

// Reader provides a random reader that can reset during testing.
var Reader io.Reader

// Int63n returns a int64 between zero and value of max, read from an io.Reader source.
func Int63n(reader io.Reader, max int64) (int64, error) {
	bi, err := rand.Int(reader, big.NewInt(max))
	if err != nil {
		return 0, fmt.Errorf("failed to read random value, %w", err)
	}

	return bi.Int64(), nil
}

// CryptoRandInt63n returns a random int64 between zero and value of max
// obtained from the crypto rand source.
func CryptoRandInt63n(max int64) (int64, error) {
	return Int63n(Reader, max)
}

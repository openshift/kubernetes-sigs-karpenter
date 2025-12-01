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

var floatMaxBigInt = big.NewInt(1 << 53)

// Float64 returns a float64 read from an io.Reader source. The returned float will be between [0.0, 1.0).
func Float64(reader io.Reader) (float64, error) {
	bi, err := rand.Int(reader, floatMaxBigInt)
	if err != nil {
		return 0, fmt.Errorf("failed to read random value, %v", err)
	}

	return float64(bi.Int64()) / (1 << 53), nil
}

// CryptoRandFloat64 returns a random float64 obtained from the crypto rand
// source.
func CryptoRandFloat64() (float64, error) {
	return Float64(Reader)
}

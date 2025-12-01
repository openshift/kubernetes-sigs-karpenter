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

// Copyright 2013-2022 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package properties

import (
	"fmt"
	"math"
)

// make this a var to overwrite it in a test
var is32Bit = ^uint(0) == math.MaxUint32

// intRangeCheck checks if the value fits into the int type and
// panics if it does not.
func intRangeCheck(key string, v int64) int {
	if is32Bit && (v < math.MinInt32 || v > math.MaxInt32) {
		panic(fmt.Sprintf("Value %d for key %s out of range", v, key))
	}
	return int(v)
}

// uintRangeCheck checks if the value fits into the uint type and
// panics if it does not.
func uintRangeCheck(key string, v uint64) uint {
	if is32Bit && v > math.MaxUint32 {
		panic(fmt.Sprintf("Value %d for key %s out of range", v, key))
	}
	return uint(v)
}

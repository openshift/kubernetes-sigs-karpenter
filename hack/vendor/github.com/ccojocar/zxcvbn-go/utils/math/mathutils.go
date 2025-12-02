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

package zxcvbnmath

import "math"

/*
NChoseK http://blog.plover.com/math/choose.html
I am surprised that I have to define these. . . Maybe i just didn't look hard enough for a lib.
*/
func NChoseK(n, k float64) float64 {
	if k > n {
		return 0
	} else if k == 0 {
		return 1
	}

	var r float64 = 1

	for d := float64(1); d <= k; d++ {
		r *= n
		r /= d
		n--
	}

	return r
}

// Round a number
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

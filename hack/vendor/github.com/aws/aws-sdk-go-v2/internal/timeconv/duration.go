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

package timeconv

import "time"

// FloatSecondsDur converts a fractional seconds to duration.
func FloatSecondsDur(v float64) time.Duration {
	return time.Duration(v * float64(time.Second))
}

// DurSecondsFloat converts a duration into fractional seconds.
func DurSecondsFloat(d time.Duration) float64 {
	return float64(d) / float64(time.Second)
}

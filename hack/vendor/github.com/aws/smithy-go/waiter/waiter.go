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

package waiter

import (
	"fmt"
	"math"
	"time"

	"github.com/aws/smithy-go/rand"
)

// ComputeDelay computes delay between waiter attempts. The function takes in a current attempt count,
// minimum delay, maximum delay, and remaining wait time for waiter as input. The inputs minDelay and maxDelay
// must always be greater than 0, along with minDelay lesser than or equal to maxDelay.
//
// Returns the computed delay and if next attempt count is possible within the given input time constraints.
// Note that the zeroth attempt results in no delay.
func ComputeDelay(attempt int64, minDelay, maxDelay, remainingTime time.Duration) (delay time.Duration, err error) {
	// zeroth attempt, no delay
	if attempt <= 0 {
		return 0, nil
	}

	// remainingTime is zero or less, no delay
	if remainingTime <= 0 {
		return 0, nil
	}

	// validate min delay is greater than 0
	if minDelay == 0 {
		return 0, fmt.Errorf("minDelay must be greater than zero when computing Delay")
	}

	// validate max delay is greater than 0
	if maxDelay == 0 {
		return 0, fmt.Errorf("maxDelay must be greater than zero when computing Delay")
	}

	// Get attempt ceiling to prevent integer overflow.
	attemptCeiling := (math.Log(float64(maxDelay/minDelay)) / math.Log(2)) + 1

	if attempt > int64(attemptCeiling) {
		delay = maxDelay
	} else {
		// Compute exponential delay based on attempt.
		ri := 1 << uint64(attempt-1)
		// compute delay
		delay = minDelay * time.Duration(ri)
	}

	if delay != minDelay {
		// randomize to get jitter between min delay and delay value
		d, err := rand.CryptoRandInt63n(int64(delay - minDelay))
		if err != nil {
			return 0, fmt.Errorf("error computing retry jitter, %w", err)
		}

		delay = time.Duration(d) + minDelay
	}

	// check if this is the last attempt possible and compute delay accordingly
	if remainingTime-delay <= minDelay {
		delay = remainingTime - minDelay
	}

	return delay, nil
}

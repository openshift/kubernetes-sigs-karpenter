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

package retry

import (
	"math"
	"time"

	"github.com/aws/aws-sdk-go-v2/internal/rand"
	"github.com/aws/aws-sdk-go-v2/internal/timeconv"
)

// ExponentialJitterBackoff provides backoff delays with jitter based on the
// number of attempts.
type ExponentialJitterBackoff struct {
	maxBackoff time.Duration
	// precomputed number of attempts needed to reach max backoff.
	maxBackoffAttempts float64

	randFloat64 func() (float64, error)
}

// NewExponentialJitterBackoff returns an ExponentialJitterBackoff configured
// for the max backoff.
func NewExponentialJitterBackoff(maxBackoff time.Duration) *ExponentialJitterBackoff {
	return &ExponentialJitterBackoff{
		maxBackoff: maxBackoff,
		maxBackoffAttempts: math.Log2(
			float64(maxBackoff) / float64(time.Second)),
		randFloat64: rand.CryptoRandFloat64,
	}
}

// BackoffDelay returns the duration to wait before the next attempt should be
// made. Returns an error if unable get a duration.
func (j *ExponentialJitterBackoff) BackoffDelay(attempt int, err error) (time.Duration, error) {
	if attempt > int(j.maxBackoffAttempts) {
		return j.maxBackoff, nil
	}

	b, err := j.randFloat64()
	if err != nil {
		return 0, err
	}

	// [0.0, 1.0) * 2 ^ attempts
	ri := int64(1 << uint64(attempt))
	delaySeconds := b * float64(ri)

	return timeconv.FloatSecondsDur(delaySeconds), nil
}

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

package sdk

import (
	"context"
	"time"
)

func init() {
	NowTime = time.Now
	Sleep = time.Sleep
	SleepWithContext = sleepWithContext
}

// NowTime is a value for getting the current time. This value can be overridden
// for testing mocking out current time.
var NowTime func() time.Time

// Sleep is a value for sleeping for a duration. This value can be overridden
// for testing and mocking out sleep duration.
var Sleep func(time.Duration)

// SleepWithContext will wait for the timer duration to expire, or the context
// is canceled. Which ever happens first. If the context is canceled the Context's
// error will be returned.
//
// This value can be overridden for testing and mocking out sleep duration.
var SleepWithContext func(context.Context, time.Duration) error

// sleepWithContext will wait for the timer duration to expire, or the context
// is canceled. Which ever happens first. If the context is canceled the
// Context's error will be returned.
func sleepWithContext(ctx context.Context, dur time.Duration) error {
	t := time.NewTimer(dur)
	defer t.Stop()

	select {
	case <-t.C:
		break
	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}

// noOpSleepWithContext does nothing, returns immediately.
func noOpSleepWithContext(context.Context, time.Duration) error {
	return nil
}

func noOpSleep(time.Duration) {}

// TestingUseNopSleep is a utility for disabling sleep across the SDK for
// testing.
func TestingUseNopSleep() func() {
	SleepWithContext = noOpSleepWithContext
	Sleep = noOpSleep

	return func() {
		SleepWithContext = sleepWithContext
		Sleep = time.Sleep
	}
}

// TestingUseReferenceTime is a utility for swapping the time function across the SDK to return a specific reference time
// for testing purposes.
func TestingUseReferenceTime(referenceTime time.Time) func() {
	NowTime = func() time.Time {
		return referenceTime
	}
	return func() {
		NowTime = time.Now
	}
}

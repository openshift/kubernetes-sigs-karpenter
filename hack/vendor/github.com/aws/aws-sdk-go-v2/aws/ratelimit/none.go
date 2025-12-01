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

package ratelimit

import "context"

// None implements a no-op rate limiter which effectively disables client-side
// rate limiting (also known as "retry quotas").
//
// GetToken does nothing and always returns a nil error. The returned
// token-release function does nothing, and always returns a nil error.
//
// AddTokens does nothing and always returns a nil error.
var None = &none{}

type none struct{}

func (*none) GetToken(ctx context.Context, cost uint) (func() error, error) {
	return func() error { return nil }, nil
}

func (*none) AddTokens(v uint) error { return nil }

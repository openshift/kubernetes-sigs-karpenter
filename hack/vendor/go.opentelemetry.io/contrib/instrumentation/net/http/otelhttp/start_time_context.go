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

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelhttp // import "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

import (
	"context"
	"time"
)

type startTimeContextKeyType int

const startTimeContextKey startTimeContextKeyType = 0

// ContextWithStartTime returns a new context with the provided start time. The
// start time will be used for metrics and traces emitted by the
// instrumentation. Only one labeller can be injected into the context.
// Injecting it multiple times will override the previous calls.
func ContextWithStartTime(parent context.Context, start time.Time) context.Context {
	return context.WithValue(parent, startTimeContextKey, start)
}

// StartTimeFromContext retrieves a time.Time from the provided context if one
// is available. If no start time was found in the provided context, a new,
// zero start time is returned and the second return value is false.
func StartTimeFromContext(ctx context.Context) time.Time {
	t, _ := ctx.Value(startTimeContextKey).(time.Time)
	return t
}

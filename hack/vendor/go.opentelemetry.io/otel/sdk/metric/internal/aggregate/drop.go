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

package aggregate // import "go.opentelemetry.io/otel/sdk/metric/internal/aggregate"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric/exemplar"
)

// dropReservoir returns a [FilteredReservoir] that drops all measurements it is offered.
func dropReservoir[N int64 | float64](attribute.Set) FilteredExemplarReservoir[N] {
	return &dropRes[N]{}
}

type dropRes[N int64 | float64] struct{}

// Offer does nothing, all measurements offered will be dropped.
func (r *dropRes[N]) Offer(context.Context, N, []attribute.KeyValue) {}

// Collect resets dest. No exemplars will ever be returned.
func (r *dropRes[N]) Collect(dest *[]exemplar.Exemplar) {
	clear(*dest) // Erase elements to let GC collect objects
	*dest = (*dest)[:0]
}

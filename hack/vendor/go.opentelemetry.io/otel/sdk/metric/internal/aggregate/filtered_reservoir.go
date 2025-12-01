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
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric/exemplar"
)

// FilteredExemplarReservoir wraps a [exemplar.Reservoir] with a filter.
type FilteredExemplarReservoir[N int64 | float64] interface {
	// Offer accepts the parameters associated with a measurement. The
	// parameters will be stored as an exemplar if the filter decides to
	// sample the measurement.
	//
	// The passed ctx needs to contain any baggage or span that were active
	// when the measurement was made. This information may be used by the
	// Reservoir in making a sampling decision.
	Offer(ctx context.Context, val N, attr []attribute.KeyValue)
	// Collect returns all the held exemplars in the reservoir.
	Collect(dest *[]exemplar.Exemplar)
}

// filteredExemplarReservoir handles the pre-sampled exemplar of measurements made.
type filteredExemplarReservoir[N int64 | float64] struct {
	filter    exemplar.Filter
	reservoir exemplar.Reservoir
}

// NewFilteredExemplarReservoir creates a [FilteredExemplarReservoir] which only offers values
// that are allowed by the filter.
func NewFilteredExemplarReservoir[N int64 | float64](
	f exemplar.Filter,
	r exemplar.Reservoir,
) FilteredExemplarReservoir[N] {
	return &filteredExemplarReservoir[N]{
		filter:    f,
		reservoir: r,
	}
}

func (f *filteredExemplarReservoir[N]) Offer(ctx context.Context, val N, attr []attribute.KeyValue) {
	if f.filter(ctx) {
		// only record the current time if we are sampling this measurement.
		f.reservoir.Offer(ctx, time.Now(), exemplar.NewValue(val), attr)
	}
}

func (f *filteredExemplarReservoir[N]) Collect(dest *[]exemplar.Exemplar) { f.reservoir.Collect(dest) }

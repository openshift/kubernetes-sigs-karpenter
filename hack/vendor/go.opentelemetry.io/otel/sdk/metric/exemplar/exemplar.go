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

package exemplar // import "go.opentelemetry.io/otel/sdk/metric/exemplar"

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
)

// Exemplar is a measurement sampled from a timeseries providing a typical
// example.
type Exemplar struct {
	// FilteredAttributes are the attributes recorded with the measurement but
	// filtered out of the timeseries' aggregated data.
	FilteredAttributes []attribute.KeyValue
	// Time is the time when the measurement was recorded.
	Time time.Time
	// Value is the measured value.
	Value Value
	// SpanID is the ID of the span that was active during the measurement. If
	// no span was active or the span was not sampled this will be empty.
	SpanID []byte `json:",omitempty"`
	// TraceID is the ID of the trace the active span belonged to during the
	// measurement. If no span was active or the span was not sampled this will
	// be empty.
	TraceID []byte `json:",omitempty"`
}

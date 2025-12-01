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
	"context"

	"github.com/aws/smithy-go/metrics"
	"github.com/aws/smithy-go/middleware"
)

type attemptMetrics struct {
	Attempts metrics.Int64Counter
	Errors   metrics.Int64Counter

	AttemptDuration metrics.Float64Histogram
}

func newAttemptMetrics(meter metrics.Meter) (*attemptMetrics, error) {
	m := &attemptMetrics{}
	var err error

	m.Attempts, err = meter.Int64Counter("client.call.attempts", func(o *metrics.InstrumentOptions) {
		o.UnitLabel = "{attempt}"
		o.Description = "The number of attempts for an individual operation"
	})
	if err != nil {
		return nil, err
	}
	m.Errors, err = meter.Int64Counter("client.call.errors", func(o *metrics.InstrumentOptions) {
		o.UnitLabel = "{error}"
		o.Description = "The number of errors for an operation"
	})
	if err != nil {
		return nil, err
	}
	m.AttemptDuration, err = meter.Float64Histogram("client.call.attempt_duration", func(o *metrics.InstrumentOptions) {
		o.UnitLabel = "s"
		o.Description = "The time it takes to connect to the service, send the request, and get back HTTP status code and headers (including time queued waiting to be sent)"
	})
	if err != nil {
		return nil, err
	}

	return m, nil
}

func withOperationMetadata(ctx context.Context) metrics.RecordMetricOption {
	return func(o *metrics.RecordMetricOptions) {
		o.Properties.Set("rpc.service", middleware.GetServiceID(ctx))
		o.Properties.Set("rpc.method", middleware.GetOperationName(ctx))
	}
}

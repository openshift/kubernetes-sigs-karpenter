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

package metrics

import "context"

// NopMeterProvider is a no-op metrics implementation.
type NopMeterProvider struct{}

var _ MeterProvider = (*NopMeterProvider)(nil)

// Meter returns a meter which creates no-op instruments.
func (NopMeterProvider) Meter(string, ...MeterOption) Meter {
	return nopMeter{}
}

type nopMeter struct{}

var _ Meter = (*nopMeter)(nil)

func (nopMeter) Int64Counter(string, ...InstrumentOption) (Int64Counter, error) {
	return nopInstrument[int64]{}, nil
}
func (nopMeter) Int64UpDownCounter(string, ...InstrumentOption) (Int64UpDownCounter, error) {
	return nopInstrument[int64]{}, nil
}
func (nopMeter) Int64Gauge(string, ...InstrumentOption) (Int64Gauge, error) {
	return nopInstrument[int64]{}, nil
}
func (nopMeter) Int64Histogram(string, ...InstrumentOption) (Int64Histogram, error) {
	return nopInstrument[int64]{}, nil
}
func (nopMeter) Int64AsyncCounter(string, Int64Callback, ...InstrumentOption) (AsyncInstrument, error) {
	return nopInstrument[int64]{}, nil
}
func (nopMeter) Int64AsyncUpDownCounter(string, Int64Callback, ...InstrumentOption) (AsyncInstrument, error) {
	return nopInstrument[int64]{}, nil
}
func (nopMeter) Int64AsyncGauge(string, Int64Callback, ...InstrumentOption) (AsyncInstrument, error) {
	return nopInstrument[int64]{}, nil
}
func (nopMeter) Float64Counter(string, ...InstrumentOption) (Float64Counter, error) {
	return nopInstrument[float64]{}, nil
}
func (nopMeter) Float64UpDownCounter(string, ...InstrumentOption) (Float64UpDownCounter, error) {
	return nopInstrument[float64]{}, nil
}
func (nopMeter) Float64Gauge(string, ...InstrumentOption) (Float64Gauge, error) {
	return nopInstrument[float64]{}, nil
}
func (nopMeter) Float64Histogram(string, ...InstrumentOption) (Float64Histogram, error) {
	return nopInstrument[float64]{}, nil
}
func (nopMeter) Float64AsyncCounter(string, Float64Callback, ...InstrumentOption) (AsyncInstrument, error) {
	return nopInstrument[float64]{}, nil
}
func (nopMeter) Float64AsyncUpDownCounter(string, Float64Callback, ...InstrumentOption) (AsyncInstrument, error) {
	return nopInstrument[float64]{}, nil
}
func (nopMeter) Float64AsyncGauge(string, Float64Callback, ...InstrumentOption) (AsyncInstrument, error) {
	return nopInstrument[float64]{}, nil
}

type nopInstrument[N any] struct{}

func (nopInstrument[N]) Add(context.Context, N, ...RecordMetricOption)    {}
func (nopInstrument[N]) Sample(context.Context, N, ...RecordMetricOption) {}
func (nopInstrument[N]) Record(context.Context, N, ...RecordMetricOption) {}
func (nopInstrument[_]) Stop()                                            {}

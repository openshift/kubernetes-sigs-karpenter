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

package opentracing

type registeredTracer struct {
	tracer       Tracer
	isRegistered bool
}

var (
	globalTracer = registeredTracer{NoopTracer{}, false}
)

// SetGlobalTracer sets the [singleton] opentracing.Tracer returned by
// GlobalTracer(). Those who use GlobalTracer (rather than directly manage an
// opentracing.Tracer instance) should call SetGlobalTracer as early as
// possible in main(), prior to calling the `StartSpan` global func below.
// Prior to calling `SetGlobalTracer`, any Spans started via the `StartSpan`
// (etc) globals are noops.
func SetGlobalTracer(tracer Tracer) {
	globalTracer = registeredTracer{tracer, true}
}

// GlobalTracer returns the global singleton `Tracer` implementation.
// Before `SetGlobalTracer()` is called, the `GlobalTracer()` is a noop
// implementation that drops all data handed to it.
func GlobalTracer() Tracer {
	return globalTracer.tracer
}

// StartSpan defers to `Tracer.StartSpan`. See `GlobalTracer()`.
func StartSpan(operationName string, opts ...StartSpanOption) Span {
	return globalTracer.tracer.StartSpan(operationName, opts...)
}

// InitGlobalTracer is deprecated. Please use SetGlobalTracer.
func InitGlobalTracer(tracer Tracer) {
	SetGlobalTracer(tracer)
}

// IsGlobalTracerRegistered returns a `bool` to indicate if a tracer has been globally registered
func IsGlobalTracerRegistered() bool {
	return globalTracer.isRegistered
}

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

import "github.com/opentracing/opentracing-go/log"

// A NoopTracer is a trivial, minimum overhead implementation of Tracer
// for which all operations are no-ops.
//
// The primary use of this implementation is in libraries, such as RPC
// frameworks, that make tracing an optional feature controlled by the
// end user. A no-op implementation allows said libraries to use it
// as the default Tracer and to write instrumentation that does
// not need to keep checking if the tracer instance is nil.
//
// For the same reason, the NoopTracer is the default "global" tracer
// (see GlobalTracer and SetGlobalTracer functions).
//
// WARNING: NoopTracer does not support baggage propagation.
type NoopTracer struct{}

type noopSpan struct{}
type noopSpanContext struct{}

var (
	defaultNoopSpanContext SpanContext = noopSpanContext{}
	defaultNoopSpan        Span        = noopSpan{}
	defaultNoopTracer      Tracer      = NoopTracer{}
)

const (
	emptyString = ""
)

// noopSpanContext:
func (n noopSpanContext) ForeachBaggageItem(handler func(k, v string) bool) {}

// noopSpan:
func (n noopSpan) Context() SpanContext                                  { return defaultNoopSpanContext }
func (n noopSpan) SetBaggageItem(key, val string) Span                   { return n }
func (n noopSpan) BaggageItem(key string) string                         { return emptyString }
func (n noopSpan) SetTag(key string, value interface{}) Span             { return n }
func (n noopSpan) LogFields(fields ...log.Field)                         {}
func (n noopSpan) LogKV(keyVals ...interface{})                          {}
func (n noopSpan) Finish()                                               {}
func (n noopSpan) FinishWithOptions(opts FinishOptions)                  {}
func (n noopSpan) SetOperationName(operationName string) Span            { return n }
func (n noopSpan) Tracer() Tracer                                        { return defaultNoopTracer }
func (n noopSpan) LogEvent(event string)                                 {}
func (n noopSpan) LogEventWithPayload(event string, payload interface{}) {}
func (n noopSpan) Log(data LogData)                                      {}

// StartSpan belongs to the Tracer interface.
func (n NoopTracer) StartSpan(operationName string, opts ...StartSpanOption) Span {
	return defaultNoopSpan
}

// Inject belongs to the Tracer interface.
func (n NoopTracer) Inject(sp SpanContext, format interface{}, carrier interface{}) error {
	return nil
}

// Extract belongs to the Tracer interface.
func (n NoopTracer) Extract(format interface{}, carrier interface{}) (SpanContext, error) {
	return nil, ErrSpanContextNotFound
}

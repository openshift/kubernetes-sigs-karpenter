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

package tracing

import "context"

// NopTracerProvider is a no-op tracing implementation.
type NopTracerProvider struct{}

var _ TracerProvider = (*NopTracerProvider)(nil)

// Tracer returns a tracer which creates no-op spans.
func (NopTracerProvider) Tracer(string, ...TracerOption) Tracer {
	return nopTracer{}
}

type nopTracer struct{}

var _ Tracer = (*nopTracer)(nil)

func (nopTracer) StartSpan(ctx context.Context, name string, opts ...SpanOption) (context.Context, Span) {
	return ctx, nopSpan{}
}

type nopSpan struct{}

var _ Span = (*nopSpan)(nil)

func (nopSpan) Name() string                    { return "" }
func (nopSpan) Context() SpanContext            { return SpanContext{} }
func (nopSpan) AddEvent(string, ...EventOption) {}
func (nopSpan) SetProperty(any, any)            {}
func (nopSpan) SetStatus(SpanStatus)            {}
func (nopSpan) End()                            {}

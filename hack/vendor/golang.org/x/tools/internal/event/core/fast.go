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

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"context"

	"golang.org/x/tools/internal/event/keys"
	"golang.org/x/tools/internal/event/label"
)

// Log1 takes a message and one label delivers a log event to the exporter.
// It is a customized version of Print that is faster and does no allocation.
func Log1(ctx context.Context, message string, t1 label.Label) {
	Export(ctx, MakeEvent([3]label.Label{
		keys.Msg.Of(message),
		t1,
	}, nil))
}

// Log2 takes a message and two labels and delivers a log event to the exporter.
// It is a customized version of Print that is faster and does no allocation.
func Log2(ctx context.Context, message string, t1 label.Label, t2 label.Label) {
	Export(ctx, MakeEvent([3]label.Label{
		keys.Msg.Of(message),
		t1,
		t2,
	}, nil))
}

// Metric1 sends a label event to the exporter with the supplied labels.
func Metric1(ctx context.Context, t1 label.Label) context.Context {
	return Export(ctx, MakeEvent([3]label.Label{
		keys.Metric.New(),
		t1,
	}, nil))
}

// Metric2 sends a label event to the exporter with the supplied labels.
func Metric2(ctx context.Context, t1, t2 label.Label) context.Context {
	return Export(ctx, MakeEvent([3]label.Label{
		keys.Metric.New(),
		t1,
		t2,
	}, nil))
}

// Start1 sends a span start event with the supplied label list to the exporter.
// It also returns a function that will end the span, which should normally be
// deferred.
func Start1(ctx context.Context, name string, t1 label.Label) (context.Context, func()) {
	return ExportPair(ctx,
		MakeEvent([3]label.Label{
			keys.Start.Of(name),
			t1,
		}, nil),
		MakeEvent([3]label.Label{
			keys.End.New(),
		}, nil))
}

// Start2 sends a span start event with the supplied label list to the exporter.
// It also returns a function that will end the span, which should normally be
// deferred.
func Start2(ctx context.Context, name string, t1, t2 label.Label) (context.Context, func()) {
	return ExportPair(ctx,
		MakeEvent([3]label.Label{
			keys.Start.Of(name),
			t1,
			t2,
		}, nil),
		MakeEvent([3]label.Label{
			keys.End.New(),
		}, nil))
}

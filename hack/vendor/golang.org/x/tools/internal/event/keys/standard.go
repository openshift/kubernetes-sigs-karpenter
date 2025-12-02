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

// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package keys

var (
	// Msg is a key used to add message strings to label lists.
	Msg = NewString("message", "a readable message")
	// Label is a key used to indicate an event adds labels to the context.
	Label = NewTag("label", "a label context marker")
	// Start is used for things like traces that have a name.
	Start = NewString("start", "span start")
	// Metric is a key used to indicate an event records metrics.
	End = NewTag("end", "a span end marker")
	// Metric is a key used to indicate an event records metrics.
	Detach = NewTag("detach", "a span detach marker")
	// Err is a key used to add error values to label lists.
	Err = NewError("error", "an error that occurred")
	// Metric is a key used to indicate an event records metrics.
	Metric = NewTag("metric", "a metric event marker")
)

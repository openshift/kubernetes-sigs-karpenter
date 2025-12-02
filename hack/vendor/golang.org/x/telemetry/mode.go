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

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package telemetry

import "golang.org/x/telemetry/internal/telemetry"

// Mode returns the current telemetry mode.
//
// The telemetry mode is a global value that controls both the local collection
// and uploading of telemetry data. Possible mode values are:
//   - "on":    both collection and uploading is enabled
//   - "local": collection is enabled, but uploading is disabled
//   - "off":   both collection and uploading are disabled
//
// When mode is "on", or "local", telemetry data is written to the local file
// system and may be inspected with the [gotelemetry] command.
//
// If an error occurs while reading the telemetry mode from the file system,
// Mode returns the default value "local".
//
// [gotelemetry]: https://pkg.go.dev/golang.org/x/telemetry/cmd/gotelemetry
func Mode() string {
	mode, _ := telemetry.Default.Mode()
	return mode
}

// SetMode sets the global telemetry mode to the given value.
//
// See the documentation of [Mode] for a description of the supported mode
// values.
//
// An error is returned if the provided mode value is invalid, or if an error
// occurs while persisting the mode value to the file system.
func SetMode(mode string) error {
	return telemetry.Default.SetMode(mode)
}

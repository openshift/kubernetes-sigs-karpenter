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

// Package timeformat defines an Analyzer that checks for the use
// of time.Format or time.Parse calls with a bad format.
//
// # Analyzer timeformat
//
// timeformat: check for calls of (time.Time).Format or time.Parse with 2006-02-01
//
// The timeformat checker looks for time formats with the 2006-02-01 (yyyy-dd-mm)
// format. Internationally, "yyyy-dd-mm" does not occur in common calendar date
// standards, and so it is more likely that 2006-01-02 (yyyy-mm-dd) was intended.
package timeformat

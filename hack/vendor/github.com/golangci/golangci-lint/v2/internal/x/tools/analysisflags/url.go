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

package analysisflags

import (
	"fmt"
	"net/url"

	"golang.org/x/tools/go/analysis"
)

// ResolveURL resolves the URL field for a Diagnostic from an Analyzer
// and returns the URL. See Diagnostic.URL for details.
func ResolveURL(a *analysis.Analyzer, d analysis.Diagnostic) (string, error) {
	if d.URL == "" && d.Category == "" && a.URL == "" {
		return "", nil // do nothing
	}
	raw := d.URL
	if d.URL == "" && d.Category != "" {
		raw = "#" + d.Category
	}
	u, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("invalid Diagnostic.URL %q: %s", raw, err)
	}
	base, err := url.Parse(a.URL)
	if err != nil {
		return "", fmt.Errorf("invalid Analyzer.URL %q: %s", a.URL, err)
	}
	return base.ResolveReference(u).String(), nil
}

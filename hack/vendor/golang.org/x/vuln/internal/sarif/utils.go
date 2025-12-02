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

package sarif

import (
	"strings"

	"golang.org/x/vuln/internal/govulncheck"
)

func choose(s1, s2 string, cond bool) string {
	if cond {
		return s1
	}
	return s2
}

func list(elems []string) string {
	l := len(elems)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return elems[0]
	}

	cList := strings.Join(elems[:l-1], ", ")
	return cList + choose("", ",", l == 2) + " and " + elems[l-1]
}

// symbol is simplified adaptation of internal/scan/symbol.
func symbol(fr *govulncheck.Frame) string {
	if fr.Function == "" {
		return ""
	}
	sym := strings.Split(fr.Function, "$")[0]
	if fr.Receiver != "" {
		sym = fr.Receiver + "." + sym
	}
	if fr.Package != "" {
		sym = fr.Package + "." + sym
	}
	return sym
}

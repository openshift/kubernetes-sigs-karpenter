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

package packages

import (
	"fmt"
	"strings"
)

var modes = [...]struct {
	mode LoadMode
	name string
}{
	{NeedName, "NeedName"},
	{NeedFiles, "NeedFiles"},
	{NeedCompiledGoFiles, "NeedCompiledGoFiles"},
	{NeedImports, "NeedImports"},
	{NeedDeps, "NeedDeps"},
	{NeedExportFile, "NeedExportFile"},
	{NeedTypes, "NeedTypes"},
	{NeedSyntax, "NeedSyntax"},
	{NeedTypesInfo, "NeedTypesInfo"},
	{NeedTypesSizes, "NeedTypesSizes"},
	{NeedForTest, "NeedForTest"},
	{NeedModule, "NeedModule"},
	{NeedEmbedFiles, "NeedEmbedFiles"},
	{NeedEmbedPatterns, "NeedEmbedPatterns"},
	{NeedTarget, "NeedTarget"},
}

func (mode LoadMode) String() string {
	if mode == 0 {
		return "LoadMode(0)"
	}
	var out []string
	// named bits
	for _, item := range modes {
		if (mode & item.mode) != 0 {
			mode ^= item.mode
			out = append(out, item.name)
		}
	}
	// unnamed residue
	if mode != 0 {
		if out == nil {
			return fmt.Sprintf("LoadMode(%#x)", int(mode))
		}
		out = append(out, fmt.Sprintf("%#x", int(mode)))
	}
	if len(out) == 1 {
		return out[0]
	}
	return "(" + strings.Join(out, "|") + ")"
}

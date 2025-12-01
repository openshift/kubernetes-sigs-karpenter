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

package analysisutil

import (
	"go/types"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// RemoVendor removes vendoring information from import path.
func RemoveVendor(path string) string {
	i := strings.Index(path, "vendor/")
	if i >= 0 {
		return path[i+len("vendor/"):]
	}
	return path
}

// LookupFromImports finds an object from import paths.
func LookupFromImports(imports []*types.Package, path, name string) types.Object {
	path = RemoveVendor(path)
	for i := range imports {
		if path == RemoveVendor(imports[i].Path()) {
			return imports[i].Scope().Lookup(name)
		}
	}
	return nil
}

// Imported returns true when the given pass imports the pkg.
func Imported(pkgPath string, pass *analysis.Pass) bool {
	fs := pass.Files
	if len(fs) == 0 {
		return false
	}
	for _, f := range fs {
		for _, i := range f.Imports {
			path, err := strconv.Unquote(i.Path.Value)
			if err != nil {
				continue
			}
			if RemoveVendor(path) == pkgPath {
				return true
			}
		}
	}
	return false
}

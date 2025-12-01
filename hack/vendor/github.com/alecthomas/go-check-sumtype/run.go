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

package gochecksumtype

import "golang.org/x/tools/go/packages"

// Run sumtype checking on the given packages.
func Run(pkgs []*packages.Package, config Config) []error {
	var errs []error

	decls, err := findSumTypeDecls(pkgs)
	if err != nil {
		return []error{err}
	}

	defs, defErrs := findSumTypeDefs(decls)
	errs = append(errs, defErrs...)
	if len(defs) == 0 {
		return errs
	}

	for _, pkg := range pkgs {
		if pkgErrs := check(pkg, defs, config); pkgErrs != nil {
			errs = append(errs, pkgErrs...)
		}
	}
	return errs
}

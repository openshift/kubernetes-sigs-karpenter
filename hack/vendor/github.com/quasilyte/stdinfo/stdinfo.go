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

package stdinfo

type Package struct {
	// Name is a package name.
	// For "encoding/json" the package name is "json".
	Name string

	// Path is a package path, like "encoding/json".
	Path string

	// Freq is a package import frequency approximation.
	// A value of -1 means "unknown".
	Freq int
}

// PathByName maps a std package name to its package path.
//
// For packages with multiple choices, like "template",
// only the more common one is accessible ("text/template" in this case).
//
// This map doesn't contain extremely rare packages either.
// Use PackageList variable if you want to construct a different mapping.
//
// It's exported as map to make it easier to re-use it in libraries
// without copying.
var PathByName = generatedPathByName

// PackagesList is a list of std packages information.
// It's sorted by a package name.
var PackagesList = generatedPackagesList

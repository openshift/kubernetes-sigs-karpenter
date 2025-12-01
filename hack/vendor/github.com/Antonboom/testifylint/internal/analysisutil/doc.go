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

// Package analysisutil contains functions common for `analyzer` and `internal/checkers` packages.
// In addition, it is intended to "lighten" these packages.
//
// If the function is common to several packages, or it makes sense to test it separately without
// "polluting" the target package with tests of private functionality, then you can put function in this package.
//
// It's important to avoid dependency on `golang.org/x/tools/go/analysis` in the helpers API.
// This makes the API "narrower" and also allows you to test functions without some "abstraction leaks".
package analysisutil

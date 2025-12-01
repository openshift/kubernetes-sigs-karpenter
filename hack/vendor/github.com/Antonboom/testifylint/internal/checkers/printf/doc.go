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

// Package printf is a patched fork of
// https://github.com/golang/tools/blob/b6235391adb3b7f8bcfc4df81055e8f023de2688/go/analysis/passes/printf/printf.go#L538
//
// Initial discussion:
// https://go-review.googlesource.com/c/tools/+/580555/comments/dfe3ef96_b1b815d5
package printf

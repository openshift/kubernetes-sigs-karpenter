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

//go:generate go run ../generate.go

// Package quickfix contains analyzes that implement code refactorings.
// None of these analyzers produce diagnostics that have to be followed.
// Most of the time, they only provide alternative ways of doing things,
// requiring users to make informed decisions.
//
// None of these analyzes should fail a build, and they are likely useless in CI as a whole.
package quickfix

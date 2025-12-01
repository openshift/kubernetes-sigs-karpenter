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

// Package stylecheck contains analyzes that enforce style rules.
// Most of the recommendations made are universally agreed upon by the wider Go community.
// Some analyzes, however, implement stricter rules that not everyone will agree with.
// In the context of Staticcheck, these analyzes are not enabled by default.
//
// For the most part it is recommended to follow the advice given by the analyzers that are enabled by default,
// but you may want to disable additional analyzes on a case by case basis.
package stylecheck

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

// Package garif defines all the Go structures required to model a SARIF log file.
// These structures were created using the JSON-schema sarif-schema-2.1.0.json of SARIF logfiles
// available at https://github.com/oasis-tcs/sarif-spec/tree/master/Schemata.
//
// The package provides constructors for all structures (see constructors.go) These constructors
// ensure that the returned structure instantiation is valid with respect to the JSON schema and
// should be used in place of plain structure instantiation.
// The root structure is LogFile.
//
// The package provides utility decorators for the most commonly used structures (see decorators.go)
package garif

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

//go:build testify_yaml_custom && !testify_yaml_fail && !testify_yaml_default

// Package yaml is an implementation of YAML functions that calls a pluggable implementation.
//
// This implementation is selected with the testify_yaml_custom build tag.
//
//	go test -tags testify_yaml_custom
//
// This implementation can be used at build time to replace the default implementation
// to avoid linking with [gopkg.in/yaml.v3].
//
// In your test package:
//
//		import assertYaml "github.com/stretchr/testify/assert/yaml"
//
//		func init() {
//			assertYaml.Unmarshal = func (in []byte, out interface{}) error {
//				// ...
//	     			return nil
//			}
//		}
package yaml

var Unmarshal func(in []byte, out interface{}) error

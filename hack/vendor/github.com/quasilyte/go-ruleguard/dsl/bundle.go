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

package dsl

// Bundle is a rules file export manifest.
type Bundle struct {
	// TODO: figure out which fields we might want to add here.
}

// ImportRules imports all rules from the bundle and prefixes them with a specified string.
//
// Empty string prefix is something like "dot import" in Go.
// Group name collisions will result in an error.
//
// Only packages that have an exported Bundle variable can be imported.
//
// Note: right now imported bundle can't import other bundles.
// This is not a fundamental limitation but rather a precaution
// measure before we understand how it should work better.
// If you need this feature, please open an issue at github.com/quasilyte/go-ruleguard.
func ImportRules(prefix string, bundle Bundle) {}

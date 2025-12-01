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

package godot

// Settings contains linter settings.
type Settings struct {
	// Which comments to check (top level declarations, top level, all).
	Scope Scope

	// Regexp for excluding particular comment lines from check.
	Exclude []string

	// Check periods at the end of sentences.
	Period bool

	// Check that first letter of each sentence is capital.
	Capital bool
}

// Scope sets which comments should be checked.
type Scope string

// List of available check scopes.
const (
	// DeclScope is for top level declaration comments.
	DeclScope Scope = "declarations"
	// TopLevelScope is for all top level comments.
	TopLevelScope Scope = "toplevel"
	// NoInlineScope is for all except inline comments.
	NoInlineScope Scope = "noinline"
	// AllScope is for all comments.
	AllScope Scope = "all"
)

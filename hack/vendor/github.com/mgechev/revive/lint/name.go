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

package lint

import "github.com/mgechev/revive/internal/rule"

// Name returns a different name if it should be different.
//
// Deprecated: Do not use this function, it will be removed in the next major release.
func Name(name string, allowlist, blocklist []string) string {
	return rule.Name(name, allowlist, blocklist, false)
}

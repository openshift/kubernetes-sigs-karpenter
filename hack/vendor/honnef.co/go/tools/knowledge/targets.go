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

package knowledge

var KnownGOOS = map[string]struct{}{
	"aix":       {},
	"android":   {},
	"darwin":    {},
	"dragonfly": {},
	"freebsd":   {},
	"hurd":      {},
	"illumos":   {},
	"ios":       {},
	"js":        {},
	"linux":     {},
	"netbsd":    {},
	"openbsd":   {},
	"plan9":     {},
	"solaris":   {},
	"wasip1":    {},
	"windows":   {},
}

var KnownGOARCH = map[string]struct{}{
	"386":      {},
	"amd64":    {},
	"arm":      {},
	"arm64":    {},
	"loong64":  {},
	"mips":     {},
	"mipsle":   {},
	"mips64":   {},
	"mips64le": {},
	"ppc64":    {},
	"ppc64le":  {},
	"riscv64":  {},
	"s390x":    {},
	"sparc64":  {},
	"wasm":     {},
}

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

package mirror

import "github.com/butuzov/mirror/internal/checker"

var (
	MaphashFunctions = []checker.Violation{
		{ // maphash.Bytes
			Targets:   checker.Bytes,
			Type:      checker.Function,
			Package:   "hash/maphash",
			Caller:    "Bytes",
			Args:      []int{1},
			AltCaller: "String",

			Generate: &checker.Generate{
				Pattern: `Bytes(maphash.MakeSeed(), $0)`,
				Returns: []string{"uint64"},
			},
		},
		{ // maphash.String
			Targets:   checker.Strings,
			Type:      checker.Function,
			Package:   "hash/maphash",
			Caller:    "String",
			Args:      []int{1},
			AltCaller: "Bytes",

			Generate: &checker.Generate{
				Pattern: `String(maphash.MakeSeed(), $0)`,
				Returns: []string{"uint64"},
			},
		},
	}

	MaphashMethods = []checker.Violation{
		{ // (*hash/maphash).Write
			Targets:   checker.Bytes,
			Type:      checker.Method,
			Package:   "hash/maphash",
			Struct:    "Hash",
			Caller:    "Write",
			Args:      []int{0},
			AltCaller: "WriteString",

			Generate: &checker.Generate{
				PreCondition: `h := maphash.Hash{}`,
				Pattern:      `Write($0)`,
				Returns:      []string{"int", "error"},
			},
		},
		{ // (*hash/maphash).WriteString
			Targets:   checker.Strings,
			Type:      checker.Method,
			Package:   "hash/maphash",
			Struct:    "Hash",
			Caller:    "WriteString",
			Args:      []int{0},
			AltCaller: "Write",

			Generate: &checker.Generate{
				PreCondition: `h := maphash.Hash{}`,
				Pattern:      `WriteString($0)`,
				Returns:      []string{"int", "error"},
			},
		},
	}
)

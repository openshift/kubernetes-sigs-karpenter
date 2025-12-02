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

package analyzer

import (
	"flag"
)

const (
	FlagNameConstRequireSingleConst = "const-require-single-const"
	FlagNameConstRequireGrouping    = "const-require-grouping"

	FlagNameImportRequireSingleImport = "import-require-single-import"
	FlagNameImportRequireGrouping     = "import-require-grouping"

	FlagNameTypeRequireSingleType = "type-require-single-type"
	FlagNameTypeRequireGrouping   = "type-require-grouping"

	FlagNameVarRequireSingleVar = "var-require-single-var"
	FlagNameVarRequireGrouping  = "var-require-grouping"
)

func Flags() flag.FlagSet {
	fs := flag.NewFlagSet(Name, flag.ExitOnError)

	fs.Bool(FlagNameConstRequireSingleConst, false, "require the use of a single global 'const' declaration only")
	fs.Bool(FlagNameConstRequireGrouping, false, "require the use of grouped global 'const' declarations")

	fs.Bool(FlagNameImportRequireSingleImport, false, "require the use of a single 'import' declaration only")
	fs.Bool(FlagNameImportRequireGrouping, false, "require the use of grouped 'import' declarations")

	fs.Bool(FlagNameTypeRequireSingleType, false, "require the use of a single global 'type' declaration only")
	fs.Bool(FlagNameTypeRequireGrouping, false, "require the use of grouped global 'type' declarations")

	fs.Bool(FlagNameVarRequireSingleVar, false, "require the use of a single global 'var' declaration only")
	fs.Bool(FlagNameVarRequireGrouping, false, "require the use of grouped global 'var' declarations")

	return *fs
}

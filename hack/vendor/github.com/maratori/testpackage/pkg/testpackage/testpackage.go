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

package testpackage

import (
	"flag"
	"regexp"
	"strings"

	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const (
	SkipRegexpFlagName    = "skip-regexp"
	SkipRegexpFlagUsage   = `regexp pattern to skip file by name. To not skip files use -skip-regexp="^$"`
	SkipRegexpFlagDefault = `(export|internal)_test\.go`
)

const (
	AllowPackagesFlagName    = "allow-packages"
	AllowPackagesFlagUsage   = `comma separated list of packages that don't end with _test that tests are allowed to be in`
	AllowPackagesFlagDefault = `main`
)

func processTestFile(pass *analysis.Pass, f *ast.File, allowedPackages []string) {
	packageName := f.Name.Name

	for _, p := range allowedPackages {
		if p == packageName {
			return
		}
	}

	if !strings.HasSuffix(packageName, "_test") {
		pass.Reportf(f.Name.Pos(), "package should be `%s_test` instead of `%s`", packageName, packageName)
	}
}

// NewAnalyzer returns Analyzer that makes you use a separate _test package.
func NewAnalyzer() *analysis.Analyzer {
	var (
		skipFileRegexp   = SkipRegexpFlagDefault
		allowPackagesStr = AllowPackagesFlagDefault
		fs               flag.FlagSet
	)

	fs.StringVar(&skipFileRegexp, SkipRegexpFlagName, skipFileRegexp, SkipRegexpFlagUsage)
	fs.StringVar(&allowPackagesStr, AllowPackagesFlagName, allowPackagesStr, AllowPackagesFlagUsage)

	return &analysis.Analyzer{
		Name:  "testpackage",
		Doc:   "linter that makes you use a separate _test package",
		Flags: fs,
		Run: func(pass *analysis.Pass) (interface{}, error) {
			allowedPackages := strings.Split(allowPackagesStr, ",")
			skipFile, err := regexp.Compile(skipFileRegexp)
			if err != nil {
				return nil, err
			}

			for _, f := range pass.Files {
				fileName := pass.Fset.Position(f.Pos()).Filename
				if !strings.HasSuffix(fileName, "_test.go") || skipFile.MatchString(fileName) {
					continue
				}

				processTestFile(pass, f, allowedPackages)
			}

			return nil, nil
		},
	}
}

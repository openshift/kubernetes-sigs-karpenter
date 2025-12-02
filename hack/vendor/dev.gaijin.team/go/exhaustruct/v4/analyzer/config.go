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
	"strings"

	"dev.gaijin.team/go/golib/e"

	"dev.gaijin.team/go/exhaustruct/v4/internal/pattern"
)

type Config struct {
	// IncludeRx is a list of regular expressions to match type names that should be
	// processed. Anonymous structs can be matched by '<anonymous>' alias.
	//
	// Each regular expression must match the full type name, including package path.
	// For example, to match type `net/http.Cookie` regular expression should be
	// `.*/http\.Cookie`, but not `http\.Cookie`.
	IncludeRx       []string     `exhaustruct:"optional"`
	includePatterns pattern.List `exhaustruct:"optional"`

	// ExcludeRx is a list of regular expressions to match type names that should be
	// excluded from processing. Anonymous structs can be matched by '<anonymous>'
	// alias.
	//
	// Has precedence over IncludeRx.
	//
	// Each regular expression must match the full type name, including package path.
	// For example, to match type `net/http.Cookie` regular expression should be
	// `.*/http\.Cookie`, but not `http\.Cookie`.
	ExcludeRx       []string     `exhaustruct:"optional"`
	excludePatterns pattern.List `exhaustruct:"optional"`

	// AllowEmpty allows empty structures, effectively excluding them from the check.
	AllowEmpty bool `exhaustruct:"optional"`

	// AllowEmptyRx is a list of regular expressions to match type names that should
	// be allowed to be empty. Anonymous structs can be matched by '<anonymous>'
	// alias.
	//
	// Each regular expression must match the full type name, including package path.
	// For example, to match type `net/http.Cookie` regular expression should be
	// `.*/http\.Cookie`, but not `http\.Cookie`.
	AllowEmptyRx       []string     `exhaustruct:"optional"`
	allowEmptyPatterns pattern.List `exhaustruct:"optional"`

	// AllowEmptyReturns allows empty structures in return statements.
	AllowEmptyReturns bool `exhaustruct:"optional"`

	// AllowEmptyDeclarations allows empty structures in variable declarations.
	AllowEmptyDeclarations bool `exhaustruct:"optional"`
}

// Prepare compiles all regular expression patterns into pattern lists for
// efficient matching.
func (c *Config) Prepare() error {
	var err error

	c.includePatterns, err = pattern.NewList(c.IncludeRx...)
	if err != nil {
		return e.NewFrom("compile include patterns", err)
	}

	c.excludePatterns, err = pattern.NewList(c.ExcludeRx...)
	if err != nil {
		return e.NewFrom("compile exclude patterns", err)
	}

	c.allowEmptyPatterns, err = pattern.NewList(c.AllowEmptyRx...)
	if err != nil {
		return e.NewFrom("compile allow empty patterns", err)
	}

	return nil
}

// stringSliceFlag implements flag.Value interface for []string fields.
type stringSliceFlag struct {
	slice *[]string
}

func (s stringSliceFlag) String() string {
	if s.slice == nil {
		return ""
	}

	return strings.Join(*s.slice, ",")
}

func (s stringSliceFlag) Set(value string) error {
	*s.slice = append(*s.slice, value)
	return nil
}

// BindToFlagSet binds the config fields to the provided flag set.
func (c *Config) BindToFlagSet(fs *flag.FlagSet) *flag.FlagSet {
	fs.Var(stringSliceFlag{&c.IncludeRx}, "include-rx",
		"Regular expression to match type names that should be processed. "+
			"Anonymous structs can be matched by '<anonymous>' alias. "+
			"Each regex must match the full type name including package path. "+
			"Example: `.*/http\\.Cookie`. Can be used multiple times.")
	fs.Var(stringSliceFlag{&c.IncludeRx}, "i", "Short form of -include-rx")

	fs.Var(stringSliceFlag{&c.ExcludeRx}, "exclude-rx",
		"Regular expression to exclude type names from processing, has precedence over -include. "+
			"Anonymous structs can be matched by '<anonymous>' alias. "+
			"Each regex must match the full type name including package path. "+
			"Example: `.*/http\\.Cookie`. Can be used multiple times.")
	fs.Var(stringSliceFlag{&c.ExcludeRx}, "e", "Short form of -exclude-rx")

	fs.BoolVar(&c.AllowEmpty, "allow-empty", c.AllowEmpty,
		"Allow empty structures, effectively excluding them from the check")

	fs.Var(stringSliceFlag{&c.AllowEmptyRx}, "allow-empty-rx",
		"Regular expression to match type names that should be allowed to be empty. "+
			"Anonymous structs can be matched by '<anonymous>' alias. "+
			"Each regex must match the full type name including package path. "+
			"Example: `.*/http\\.Cookie`. Can be used multiple times.")

	fs.BoolVar(&c.AllowEmptyReturns, "allow-empty-returns", c.AllowEmptyReturns,
		"Allow empty structures in return statements")

	fs.BoolVar(&c.AllowEmptyDeclarations, "allow-empty-declarations", c.AllowEmptyDeclarations,
		"Allow empty structures in variable declarations")

	return fs
}

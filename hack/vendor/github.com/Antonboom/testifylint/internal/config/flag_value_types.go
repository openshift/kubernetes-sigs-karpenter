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

package config

import (
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/Antonboom/testifylint/internal/checkers"
)

var (
	_ flag.Value = (*KnownCheckersValue)(nil)
	_ flag.Value = (*RegexpValue)(nil)
	_ flag.Value = (*EnumValue[checkers.SuiteExtraAssertCallMode])(nil)
)

// KnownCheckersValue implements comma separated list of testify checkers.
type KnownCheckersValue []string

func (kcv KnownCheckersValue) String() string {
	return strings.Join(kcv, ",")
}

func (kcv *KnownCheckersValue) Set(v string) error {
	chckrs := strings.Split(v, ",")
	for _, checkerName := range chckrs {
		if ok := checkers.IsKnown(checkerName); !ok {
			return fmt.Errorf("unknown checker %q", checkerName)
		}
	}

	*kcv = chckrs
	return nil
}

func (kcv KnownCheckersValue) Contains(v string) bool {
	for _, checker := range kcv {
		if checker == v {
			return true
		}
	}
	return false
}

// RegexpValue is a special wrapper for support of flag.FlagSet over regexp.Regexp.
// Original regexp is available through RegexpValue.Regexp.
type RegexpValue struct {
	*regexp.Regexp
}

func (rv RegexpValue) String() string {
	if rv.Regexp == nil {
		return ""
	}
	return rv.Regexp.String()
}

func (rv *RegexpValue) Set(v string) error {
	compiled, err := regexp.Compile(v)
	if err != nil {
		return err
	}

	rv.Regexp = compiled
	return nil
}

// EnumValue is a special type for support of flag.FlagSet over user-defined constants.
type EnumValue[EnumT comparable] struct {
	mapping map[string]EnumT
	keys    []string
	dst     *EnumT
}

// NewEnumValue takes the "enum-value-name to enum-value" mapping and a destination for the value passed through the CLI.
// Returns an EnumValue instance suitable for flag.FlagSet.Var.
func NewEnumValue[EnumT comparable](mapping map[string]EnumT, dst *EnumT) *EnumValue[EnumT] {
	keys := make([]string, 0, len(mapping))
	for k := range mapping {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return &EnumValue[EnumT]{
		mapping: mapping,
		keys:    keys,
		dst:     dst,
	}
}

func (e EnumValue[EnumT]) String() string {
	if e.dst == nil {
		return ""
	}

	for k, v := range e.mapping {
		if v == *e.dst {
			return k
		}
	}
	return ""
}

func (e *EnumValue[EnumT]) Set(s string) error {
	v, ok := e.mapping[s]
	if !ok {
		return fmt.Errorf("use one of (%v)", strings.Join(e.keys, " | "))
	}

	*e.dst = v
	return nil
}

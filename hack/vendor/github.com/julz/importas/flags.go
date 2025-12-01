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

package importas

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

var errWrongAlias = errors.New("import flag must be of form path:alias")

func flags(config *Config) flag.FlagSet {
	fs := flag.FlagSet{}
	fs.Var(&config.RequiredAlias, "alias", "required import alias in form path:alias")
	fs.BoolVar(&config.DisallowUnaliased, "no-unaliased", false, "do not allow unaliased imports of aliased packages")
	fs.BoolVar(&config.DisallowExtraAliases, "no-extra-aliases", false, "do not allow non-required aliases")
	return fs
}

type aliasList [][]string

func (v *aliasList) Set(val string) error {
	lastColon := strings.LastIndex(val, ":")
	if lastColon <= 1 {
		return errWrongAlias
	}
	*v = append(*v, []string{val[:lastColon], val[lastColon+1:]})
	return nil
}

func (v *aliasList) String() string {
	return fmt.Sprintf("%v", ([][]string)(*v))
}

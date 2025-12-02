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

package section

import (
	"fmt"
	"strings"

	"github.com/daixiang0/gci/pkg/parse"
	"github.com/daixiang0/gci/pkg/specificity"
)

type Custom struct {
	Prefix string
}

// CustomSeparator allows you to group multiple custom prefix together in the same section
// gci diff -s standard -s default -s prefix(github.com/company,gitlab.com/company,companysuffix)
const CustomSeparator = ","

const CustomType = "custom"

func (c Custom) MatchSpecificity(spec *parse.GciImports) specificity.MatchSpecificity {
	for _, prefix := range strings.Split(c.Prefix, CustomSeparator) {
		prefix = strings.TrimSpace(prefix)
		if strings.HasPrefix(spec.Path, prefix) {
			return specificity.Match{Length: len(prefix)}
		}
	}

	return specificity.MisMatch{}
}

func (c Custom) String() string {
	return fmt.Sprintf("prefix(%s)", c.Prefix)
}

func (c Custom) Type() string {
	return CustomType
}

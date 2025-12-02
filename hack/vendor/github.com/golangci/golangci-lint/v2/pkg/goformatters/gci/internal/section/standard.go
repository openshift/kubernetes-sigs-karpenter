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
	"github.com/daixiang0/gci/pkg/parse"
	"github.com/daixiang0/gci/pkg/specificity"
)

const StandardType = "standard"

type Standard struct{}

func (s Standard) MatchSpecificity(spec *parse.GciImports) specificity.MatchSpecificity {
	if isStandard(spec.Path) {
		return specificity.StandardMatch{}
	}
	return specificity.MisMatch{}
}

func (s Standard) String() string {
	return StandardType
}

func (s Standard) Type() string {
	return StandardType
}

func isStandard(pkg string) bool {
	_, ok := standardPackages[pkg]
	return ok
}

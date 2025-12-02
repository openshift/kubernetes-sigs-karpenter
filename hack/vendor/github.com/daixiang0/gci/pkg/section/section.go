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

// Section defines a part of the formatted output.
type Section interface {
	// MatchSpecificity returns how well an Import matches to this Section
	MatchSpecificity(spec *parse.GciImports) specificity.MatchSpecificity

	// String Implements the stringer interface
	String() string

	// return section type
	Type() string
}

type SectionList []Section

func (list SectionList) String() []string {
	var output []string
	for _, section := range list {
		output = append(output, section.String())
	}
	return output
}

func DefaultSections() SectionList {
	return SectionList{Standard{}, Default{}}
}

func DefaultSectionSeparators() SectionList {
	return SectionList{NewLine{}}
}

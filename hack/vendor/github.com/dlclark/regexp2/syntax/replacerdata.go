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

package syntax

import (
	"bytes"
	"errors"
)

type ReplacerData struct {
	Rep     string
	Strings []string
	Rules   []int
}

const (
	replaceSpecials     = 4
	replaceLeftPortion  = -1
	replaceRightPortion = -2
	replaceLastGroup    = -3
	replaceWholeString  = -4
)

//ErrReplacementError is a general error during parsing the replacement text
var ErrReplacementError = errors.New("Replacement pattern error.")

// NewReplacerData will populate a reusable replacer data struct based on the given replacement string
// and the capture group data from a regexp
func NewReplacerData(rep string, caps map[int]int, capsize int, capnames map[string]int, op RegexOptions) (*ReplacerData, error) {
	p := parser{
		options:  op,
		caps:     caps,
		capsize:  capsize,
		capnames: capnames,
	}
	p.setPattern(rep)
	concat, err := p.scanReplacement()
	if err != nil {
		return nil, err
	}

	if concat.t != ntConcatenate {
		panic(ErrReplacementError)
	}

	sb := &bytes.Buffer{}
	var (
		strings []string
		rules   []int
	)

	for _, child := range concat.children {
		switch child.t {
		case ntMulti:
			child.writeStrToBuf(sb)

		case ntOne:
			sb.WriteRune(child.ch)

		case ntRef:
			if sb.Len() > 0 {
				rules = append(rules, len(strings))
				strings = append(strings, sb.String())
				sb.Reset()
			}
			slot := child.m

			if len(caps) > 0 && slot >= 0 {
				slot = caps[slot]
			}

			rules = append(rules, -replaceSpecials-1-slot)

		default:
			panic(ErrReplacementError)
		}
	}

	if sb.Len() > 0 {
		rules = append(rules, len(strings))
		strings = append(strings, sb.String())
	}

	return &ReplacerData{
		Rep:     rep,
		Strings: strings,
		Rules:   rules,
	}, nil
}

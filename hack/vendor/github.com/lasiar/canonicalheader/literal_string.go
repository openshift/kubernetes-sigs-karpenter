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

package canonicalheader

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"unicode/utf8"
	"unsafe"

	"golang.org/x/tools/go/analysis"
)

type literalString struct {
	originalValue string
	quote         byte
	pos, end      token.Pos
}

func newLiteralString(basicList *ast.BasicLit) (literalString, error) {
	if basicList.Kind != token.STRING {
		return literalString{}, fmt.Errorf("%#v is not a string", basicList)
	}

	if len(basicList.Value) < 2 {
		return literalString{}, fmt.Errorf("%#v has a strange value length %q", basicList, basicList.Value)
	}

	quote := basicList.Value[0]
	switch quote {
	case '`', '"':
	default:
		return literalString{}, fmt.Errorf("%q is a strange quote", quote)
	}

	originalValue, err := strconv.Unquote(basicList.Value)
	if err != nil {
		return literalString{}, fmt.Errorf("unquote %q: %w", basicList.Value, err)
	}

	if !utf8.ValidString(originalValue) {
		return literalString{}, fmt.Errorf("%#v is not a valid utf8 string", basicList.Value)
	}

	return literalString{
		originalValue: originalValue,
		quote:         quote,
		pos:           basicList.Pos(),
		end:           basicList.End(),
	}, nil
}

func (l literalString) diagnostic(canonicalHeader string) analysis.Diagnostic {
	newText := make([]byte, 0, len(canonicalHeader)+2)
	newText = append(newText, l.quote)
	newText = append(newText, unsafe.Slice(unsafe.StringData(canonicalHeader), len(canonicalHeader))...)
	newText = append(newText, l.quote)

	return analysis.Diagnostic{
		Pos:     l.pos,
		End:     l.end,
		Message: fmt.Sprintf("non-canonical header %q, instead use: %q", l.originalValue, canonicalHeader),
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: fmt.Sprintf("should replace %q with %q", l.originalValue, canonicalHeader),
				TextEdits: []analysis.TextEdit{
					{
						Pos:     l.pos,
						End:     l.end,
						NewText: newText,
					},
				},
			},
		},
	}
}

func (l literalString) value() string {
	return l.originalValue
}

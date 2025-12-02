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

package internal

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func NewMisplacedEmbeddedFieldDiag(embeddedField *ast.Field) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos:     embeddedField.Pos(),
		Message: "embedded fields should be listed before regular fields",
	}
}

func NewMissingSpaceDiag(
	lastEmbeddedField *ast.Field,
	firstRegularField *ast.Field,
) analysis.Diagnostic {
	suggestedPos := firstRegularField.Pos()
	if firstRegularField.Doc != nil {
		suggestedPos = firstRegularField.Doc.Pos()
	}

	return analysis.Diagnostic{
		Pos:     lastEmbeddedField.Pos(),
		Message: "there must be an empty line separating embedded fields from regular fields",
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: "adding extra line separating embedded fields from regular fields",
				TextEdits: []analysis.TextEdit{
					{
						Pos:     suggestedPos,
						NewText: []byte("\n\n"),
					},
				},
			},
		},
	}
}

func NewForbiddenEmbeddedFieldDiag(forbidField *ast.SelectorExpr) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos:     forbidField.Pos(),
		Message: fmt.Sprintf("%s.%s should not be embedded", forbidField.X, forbidField.Sel.Name),
	}
}

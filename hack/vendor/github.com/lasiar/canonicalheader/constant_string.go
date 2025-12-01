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
	"go/constant"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

type constantString struct {
	originalValue,
	nameOfConst string

	pos token.Pos
	end token.Pos
}

func newConstantKey(info *types.Info, ident *ast.Ident) (constantString, error) {
	c, ok := info.ObjectOf(ident).(*types.Const)
	if !ok {
		return constantString{}, fmt.Errorf("type %T is not support", c)
	}

	return constantString{
		nameOfConst:   c.Name(),
		originalValue: constant.StringVal(c.Val()),
		pos:           ident.Pos(),
		end:           ident.End(),
	}, nil
}

func (c constantString) diagnostic(canonicalHeader string) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos: c.pos,
		End: c.end,
		Message: fmt.Sprintf(
			"const %q used as a key at http.Header, but %q is not canonical, want %q",
			c.nameOfConst,
			c.originalValue,
			canonicalHeader,
		),
	}
}

func (c constantString) value() string {
	return c.originalValue
}

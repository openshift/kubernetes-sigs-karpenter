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

package checkers

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"

	"github.com/Antonboom/testifylint/internal/analysisutil"
)

func mimicHTTPHandler(pass *analysis.Pass, fType *ast.FuncType) bool {
	httpHandlerFuncObj := analysisutil.ObjectOf(pass.Pkg, "net/http", "HandlerFunc")
	if httpHandlerFuncObj == nil {
		return false
	}

	sig, ok := httpHandlerFuncObj.Type().Underlying().(*types.Signature)
	if !ok {
		return false
	}

	if len(fType.Params.List) != sig.Params().Len() {
		return false
	}

	for i := range sig.Params().Len() {
		lhs := sig.Params().At(i).Type()
		rhs := pass.TypesInfo.TypeOf(fType.Params.List[i].Type)
		if !types.Identical(lhs, rhs) {
			return false
		}
	}
	return true
}

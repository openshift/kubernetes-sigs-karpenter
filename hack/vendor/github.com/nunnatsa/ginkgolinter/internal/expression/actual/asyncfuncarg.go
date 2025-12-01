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

package actual

import (
	gotypes "go/types"

	"github.com/nunnatsa/ginkgolinter/internal/gomegainfo"
	"github.com/nunnatsa/ginkgolinter/internal/typecheck"
)

func getAsyncFuncArg(sig *gotypes.Signature) ArgPayload {
	argType := FuncSigArgType
	if sig.Results().Len() == 1 {
		if typecheck.ImplementsError(sig.Results().At(0).Type().Underlying()) {
			argType |= ErrFuncActualArgType | ErrorTypeArgType
		}
	}

	if sig.Params().Len() > 0 {
		arg := sig.Params().At(0).Type()
		if gomegainfo.IsGomegaType(arg) && sig.Results().Len() == 0 {
			argType |= FuncSigArgType | GomegaParamArgType
		}
	}

	if sig.Results().Len() > 1 {
		argType |= FuncSigArgType | MultiRetsArgType
	}

	return &FuncSigArgPayload{argType: argType}
}

type FuncSigArgPayload struct {
	argType ArgType
}

func (f FuncSigArgPayload) ArgType() ArgType {
	return f.argType
}

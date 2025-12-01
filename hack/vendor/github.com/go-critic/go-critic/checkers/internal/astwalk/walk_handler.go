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

package astwalk

import (
	"go/ast"
)

// WalkHandler is a type to be embedded into every checker
// that uses astwalk walkers.
type WalkHandler struct {
	// SkipChilds controls whether currently analyzed
	// node childs should be traversed.
	//
	// Value is reset after each visitor invocation,
	// so there is no need to set value back to false.
	SkipChilds bool
}

// EnterFile is a default walkerEvents.EnterFile implementation
// that reports every file as accepted candidate for checking.
func (w *WalkHandler) EnterFile(_ *ast.File) bool {
	return true
}

// EnterFunc is a default walkerEvents.EnterFunc implementation
// that skips extern function (ones that do not have body).
func (w *WalkHandler) EnterFunc(decl *ast.FuncDecl) bool {
	return decl.Body != nil
}

func (w *WalkHandler) skipChilds() bool {
	v := w.SkipChilds
	w.SkipChilds = false
	return v
}

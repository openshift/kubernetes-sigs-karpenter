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
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/ast/inspector"
)

// Checker describes named checker.
type Checker interface {
	Name() string
}

// RegularChecker check assertion call presented in CallMeta form.
type RegularChecker interface {
	Checker
	Check(pass *analysis.Pass, call *CallMeta) *analysis.Diagnostic
}

// AdvancedChecker implements complex Check logic different from trivial CallMeta check.
type AdvancedChecker interface {
	Checker
	Check(pass *analysis.Pass, insp *inspector.Inspector) []analysis.Diagnostic
}

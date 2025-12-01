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

	"golang.org/x/tools/go/analysis"
)

type Slog struct {
	General
}

func (z Slog) FilterKeyAndValues(pass *analysis.Pass, keyAndValues []ast.Expr) []ast.Expr {
	// check slog.Group() constructed group slog.Attr
	// since we also check `slog.Group` so it is OK skip here
	return filterKeyAndValues(pass, keyAndValues, "Attr")
}

var _ Checker = (*Slog)(nil)

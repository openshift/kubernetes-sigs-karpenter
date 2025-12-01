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

type Zap struct {
	General
}

func (z Zap) FilterKeyAndValues(pass *analysis.Pass, keyAndValues []ast.Expr) []ast.Expr {
	// Skip any zapcore.Field we found
	// This is a strongly-typed field. Consume it and move on.
	// Actually it's go.uber.org/zap/zapcore.Field, however for simplicity
	// we don't check the import path
	return filterKeyAndValues(pass, keyAndValues, "Field")
}

var _ Checker = (*Zap)(nil)

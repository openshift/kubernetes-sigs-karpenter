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

package sa6006

import (
	"go/ast"

	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/analysis/report"
	"honnef.co/go/tools/pattern"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA6006",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	},
	Doc: &lint.RawDocumentation{
		Title: `Using io.WriteString to write \'[]byte\'`,
		Text: `Using io.WriteString to write a slice of bytes, as in

    io.WriteString(w, string(b))

is both unnecessary and inefficient. Converting from \'[]byte\' to \'string\'
has to allocate and copy the data, and we could simply use \'w.Write(b)\'
instead.`,

		Since: "2024.1",
	},
})

var Analyzer = SCAnalyzer.Analyzer

var ioWriteStringConversion = pattern.MustParse(`(CallExpr (Symbol "io.WriteString") [_ (CallExpr (Builtin "string") [arg])])`)

func run(pass *analysis.Pass) (interface{}, error) {
	fn := func(node ast.Node) {
		m, ok := code.Match(pass, ioWriteStringConversion, node)
		if !ok {
			return
		}
		if !code.IsOfStringConvertibleByteSlice(pass, m.State["arg"].(ast.Expr)) {
			return
		}
		report.Report(pass, node, "use io.Writer.Write instead of converting from []byte to string to use io.WriteString")
	}
	code.Preorder(pass, fn, (*ast.CallExpr)(nil))

	return nil, nil
}

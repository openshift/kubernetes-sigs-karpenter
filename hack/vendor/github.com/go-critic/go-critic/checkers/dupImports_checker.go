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
	"fmt"
	"go/ast"

	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "dupImport"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects multiple imports of the same package under different aliases"
	info.Before = `
import (
	"fmt"
	printing "fmt" // Imported the second time
)`
	info.After = `
import(
	"fmt"
)`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return &dupImportChecker{ctx: ctx}, nil
	})
}

type dupImportChecker struct {
	ctx *linter.CheckerContext
}

func (c *dupImportChecker) WalkFile(f *ast.File) {
	imports := make(map[string][]*ast.ImportSpec)
	for _, importDcl := range f.Imports {
		pkg := importDcl.Path.Value
		imports[pkg] = append(imports[pkg], importDcl)
	}

	for _, importList := range imports {
		if len(importList) == 1 {
			continue
		}
		c.warn(importList)
	}
}

func (c *dupImportChecker) warn(importList []*ast.ImportSpec) {
	msg := fmt.Sprintf("package is imported %d times under different aliases on lines", len(importList))
	for idx, importDcl := range importList {
		switch {
		case idx == len(importList)-1:
			msg += " and"
		case idx > 0:
			msg += ","
		}
		msg += fmt.Sprintf(" %d", c.ctx.FileSet.Position(importDcl.Pos()).Line)
	}
	for _, importDcl := range importList {
		c.ctx.Warn(importDcl, msg)
	}
}

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

//go:build !go1.21

package exhaustive

import (
	"go/ast"
	"regexp"
)

// For definition of generated file see:
// http://golang.org/s/generatedcode

var generatedCodeRe = regexp.MustCompile(`^// Code generated .* DO NOT EDIT\.$`)

func isGeneratedFile(file *ast.File) bool {
	for _, c := range file.Comments {
		for _, cc := range c.List {
			if cc.Pos() > file.Package {
				break
			}
			if generatedCodeRe.MatchString(cc.Text) {
				return true
			}
		}
	}
	return false
}

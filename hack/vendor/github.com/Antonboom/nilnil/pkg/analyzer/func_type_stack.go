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

package analyzer

import (
	"go/ast"
)

type funcTypeStack []*ast.FuncType

func (s *funcTypeStack) Push(f *ast.FuncType) {
	*s = append(*s, f)
}

func (s *funcTypeStack) Pop() *ast.FuncType {
	if len(*s) == 0 {
		return nil
	}

	last := len(*s) - 1
	f := (*s)[last]
	*s = (*s)[:last]
	return f
}

func (s *funcTypeStack) Top() *ast.FuncType {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

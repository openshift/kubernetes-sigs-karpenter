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

package golines

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dave/dst"
)

const annotationPrefix = "// __golines:shorten:"

// CreateAnnotation generates the text of a comment that will annotate long lines.
func CreateAnnotation(length int) string {
	return fmt.Sprintf(
		"%s%d",
		annotationPrefix,
		length,
	)
}

// IsAnnotation determines whether the given line is an annotation created with CreateAnnotation.
func IsAnnotation(line string) bool {
	return strings.HasPrefix(
		strings.Trim(line, " \t"),
		annotationPrefix,
	)
}

// HasAnnotation determines whether the given AST node has a line length annotation on it.
func HasAnnotation(node dst.Node) bool {
	startDecorations := node.Decorations().Start.All()
	return len(startDecorations) > 0 &&
		IsAnnotation(startDecorations[len(startDecorations)-1])
}

// HasTailAnnotation determines whether the given AST node has a line length annotation at its
// end. This is needed to catch long function declarations with inline interface definitions.
func HasTailAnnotation(node dst.Node) bool {
	endDecorations := node.Decorations().End.All()
	return len(endDecorations) > 0 &&
		IsAnnotation(endDecorations[len(endDecorations)-1])
}

// HasAnnotationRecursive determines whether the given node or one of its children has a
// golines annotation on it. It's currently implemented for function declarations, fields,
// call expressions, and selector expressions only.
func HasAnnotationRecursive(node dst.Node) bool {
	if HasAnnotation(node) {
		return true
	}

	switch n := node.(type) {
	case *dst.FuncDecl:
		if n.Type != nil && n.Type.Params != nil {
			for _, item := range n.Type.Params.List {
				if HasAnnotationRecursive(item) {
					return true
				}
			}
		}
	case *dst.Field:
		return HasTailAnnotation(n) || HasAnnotationRecursive(n.Type)
	case *dst.SelectorExpr:
		return HasAnnotation(n.Sel) || HasAnnotation(n.X)
	case *dst.CallExpr:
		if HasAnnotationRecursive(n.Fun) {
			return true
		}

		for _, arg := range n.Args {
			if HasAnnotation(arg) {
				return true
			}
		}
	case *dst.InterfaceType:
		for _, field := range n.Methods.List {
			if HasAnnotationRecursive(field) {
				return true
			}
		}
	}

	return false
}

// ParseAnnotation returns the line length encoded in a golines annotation. If none is found,
// it returns -1.
func ParseAnnotation(line string) int {
	if IsAnnotation(line) {
		components := strings.SplitN(line, ":", 3)
		val, err := strconv.Atoi(components[2])
		if err != nil {
			return -1
		}
		return val
	}
	return -1
}

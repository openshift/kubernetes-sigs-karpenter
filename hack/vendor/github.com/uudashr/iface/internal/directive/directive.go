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

package directive

import (
	"go/ast"
	"slices"
	"strings"
)

// Ignore represent a special instruction embedded in the source code.
//
// The directive can be as simple as
//
//	//iface:ignore
//
// or consist of name
//
//	//iface:ignore=unused
//
// or multiple names
//
//	//iface:ignore=unused,identical
type Ignore struct {
	Names []string
}

// ParseIgnore parse the directive from the comments.
func ParseIgnore(doc *ast.CommentGroup) *Ignore {
	if doc == nil {
		return nil
	}

	for _, comment := range doc.List {
		text := strings.TrimSpace(comment.Text)
		if text == "//iface:ignore" {
			return &Ignore{}
		}

		// parse the Names if exists
		if val, found := strings.CutPrefix(text, "//iface:ignore="); found {
			val = strings.TrimSpace(val)
			if val == "" {
				return &Ignore{}
			}

			names := strings.Split(val, ",")
			if len(names) == 0 {
				continue
			}

			for i, name := range names {
				names[i] = strings.TrimSpace(name)
			}

			if len(names) > 0 {
				return &Ignore{Names: names}
			}

			return &Ignore{}
		}
	}

	return nil
}

func (i *Ignore) hasName(name string) bool {
	return slices.Contains(i.Names, name)
}

// ShouldIgnore return true if the name should be ignored.
func (i *Ignore) ShouldIgnore(name string) bool {
	if len(i.Names) == 0 {
		return true
	}

	return i.hasName(name)
}

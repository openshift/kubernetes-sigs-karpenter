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

// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gocyclo

import (
	"go/ast"
	"strings"
)

type directives []string

func (ds directives) HasIgnore() bool {
	return ds.isPresent("ignore")
}

func (ds directives) isPresent(name string) bool {
	for _, d := range ds {
		if d == name {
			return true
		}
	}
	return false
}

func parseDirectives(doc *ast.CommentGroup) directives {
	if doc == nil {
		return directives{}
	}
	const prefix = "//gocyclo:"
	var ds directives
	for _, comment := range doc.List {
		if strings.HasPrefix(comment.Text, prefix) {
			ds = append(ds, strings.TrimSpace(strings.TrimPrefix(comment.Text, prefix)))
		}
	}
	return ds
}

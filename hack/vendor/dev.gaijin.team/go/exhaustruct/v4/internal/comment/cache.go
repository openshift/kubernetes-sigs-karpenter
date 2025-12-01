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

package comment

import (
	"go/ast"
	"go/token"
	"sync"
)

type Cache struct {
	comments map[*ast.File]ast.CommentMap
	mu       sync.RWMutex
}

// Get returns a comment map for a given file. In case if a comment map is not
// found, it creates a new one.
func (c *Cache) Get(fset *token.FileSet, f *ast.File) ast.CommentMap {
	c.mu.RLock()
	if cm, ok := c.comments[f]; ok {
		c.mu.RUnlock()
		return cm
	}
	c.mu.RUnlock()

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.comments == nil {
		c.comments = make(map[*ast.File]ast.CommentMap)
	}

	cm := ast.NewCommentMap(fset, f, f.Comments)
	c.comments[f] = cm

	return cm
}

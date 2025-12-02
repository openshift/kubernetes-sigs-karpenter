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

package structure

import (
	"go/types"
	"sync"
)

type FieldsCache struct {
	fields map[*types.Struct]Fields
	mu     sync.RWMutex
}

// Get returns a struct fields for a given type. In case if a struct fields is
// not found, it creates a new one from type definition.
func (c *FieldsCache) Get(typ *types.Struct) Fields {
	c.mu.RLock()
	fields, ok := c.fields[typ]
	c.mu.RUnlock()

	if ok {
		return fields
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.fields == nil {
		c.fields = make(map[*types.Struct]Fields)
	}

	fields = NewFields(typ)
	c.fields[typ] = fields

	return fields
}

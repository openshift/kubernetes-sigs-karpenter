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

package spec

import (
	"bytes"
	"encoding/json"
	"reflect"
	"sort"
)

// OrderSchemaItem holds a named schema (e.g. from a property of an object)
type OrderSchemaItem struct {
	Name string
	Schema
}

// OrderSchemaItems is a sortable slice of named schemas.
// The ordering is defined by the x-order schema extension.
type OrderSchemaItems []OrderSchemaItem

// MarshalJSON produces a json object with keys defined by the name schemas
// of the OrderSchemaItems slice, keeping the original order of the slice.
func (items OrderSchemaItems) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("{")
	for i := range items {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString("\"")
		buf.WriteString(items[i].Name)
		buf.WriteString("\":")
		bs, err := json.Marshal(&items[i].Schema)
		if err != nil {
			return nil, err
		}
		buf.Write(bs)
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

func (items OrderSchemaItems) Len() int      { return len(items) }
func (items OrderSchemaItems) Swap(i, j int) { items[i], items[j] = items[j], items[i] }
func (items OrderSchemaItems) Less(i, j int) (ret bool) {
	ii, oki := items[i].Extensions.GetInt("x-order")
	ij, okj := items[j].Extensions.GetInt("x-order")
	if oki {
		if okj {
			defer func() {
				if err := recover(); err != nil {
					defer func() {
						if err = recover(); err != nil {
							ret = items[i].Name < items[j].Name
						}
					}()
					ret = reflect.ValueOf(ii).String() < reflect.ValueOf(ij).String()
				}
			}()
			return ii < ij
		}
		return true
	} else if okj {
		return false
	}
	return items[i].Name < items[j].Name
}

// SchemaProperties is a map representing the properties of a Schema object.
// It knows how to transform its keys into an ordered slice.
type SchemaProperties map[string]Schema

// ToOrderedSchemaItems transforms the map of properties into a sortable slice
func (properties SchemaProperties) ToOrderedSchemaItems() OrderSchemaItems {
	items := make(OrderSchemaItems, 0, len(properties))
	for k, v := range properties {
		items = append(items, OrderSchemaItem{
			Name:   k,
			Schema: v,
		})
	}
	sort.Sort(items)
	return items
}

// MarshalJSON produces properties as json, keeping their order.
func (properties SchemaProperties) MarshalJSON() ([]byte, error) {
	if properties == nil {
		return []byte("null"), nil
	}
	return json.Marshal(properties.ToOrderedSchemaItems())
}

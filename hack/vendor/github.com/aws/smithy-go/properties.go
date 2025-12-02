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

package smithy

import "maps"

// PropertiesReader provides an interface for reading metadata from the
// underlying metadata container.
type PropertiesReader interface {
	Get(key any) any
}

// Properties provides storing and reading metadata values. Keys may be any
// comparable value type. Get and Set will panic if a key is not comparable.
//
// The zero value for a Properties instance is ready for reads/writes without
// any additional initialization.
type Properties struct {
	values map[any]any
}

// Get attempts to retrieve the value the key points to. Returns nil if the
// key was not found.
//
// Panics if key type is not comparable.
func (m *Properties) Get(key any) any {
	m.lazyInit()
	return m.values[key]
}

// Set stores the value pointed to by the key. If a value already exists at
// that key it will be replaced with the new value.
//
// Panics if the key type is not comparable.
func (m *Properties) Set(key, value any) {
	m.lazyInit()
	m.values[key] = value
}

// Has returns whether the key exists in the metadata.
//
// Panics if the key type is not comparable.
func (m *Properties) Has(key any) bool {
	m.lazyInit()
	_, ok := m.values[key]
	return ok
}

// SetAll accepts all of the given Properties into the receiver, overwriting
// any existing keys in the case of conflicts.
func (m *Properties) SetAll(other *Properties) {
	if other.values == nil {
		return
	}

	m.lazyInit()
	for k, v := range other.values {
		m.values[k] = v
	}
}

// Values returns a shallow clone of the property set's values.
func (m *Properties) Values() map[any]any {
	return maps.Clone(m.values)
}

func (m *Properties) lazyInit() {
	if m.values == nil {
		m.values = map[any]any{}
	}
}

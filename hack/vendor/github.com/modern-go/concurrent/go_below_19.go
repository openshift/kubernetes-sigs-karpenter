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

//+build !go1.9

package concurrent

import "sync"

// Map implements a thread safe map for go version below 1.9 using mutex
type Map struct {
	lock sync.RWMutex
	data map[interface{}]interface{}
}

// NewMap creates a thread safe map
func NewMap() *Map {
	return &Map{
		data: make(map[interface{}]interface{}, 32),
	}
}

// Load is same as sync.Map Load
func (m *Map) Load(key interface{}) (elem interface{}, found bool) {
	m.lock.RLock()
	elem, found = m.data[key]
	m.lock.RUnlock()
	return
}

// Load is same as sync.Map Store
func (m *Map) Store(key interface{}, elem interface{}) {
	m.lock.Lock()
	m.data[key] = elem
	m.lock.Unlock()
}

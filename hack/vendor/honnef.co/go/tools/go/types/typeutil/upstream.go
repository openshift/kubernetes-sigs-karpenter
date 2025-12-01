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

package typeutil

import (
	"go/ast"
	"go/types"
	_ "unsafe"

	"golang.org/x/tools/go/types/typeutil"
)

type MethodSetCache = typeutil.MethodSetCache
type Hasher = typeutil.Hasher

func Callee(info *types.Info, call *ast.CallExpr) types.Object {
	return typeutil.Callee(info, call)
}

func IntuitiveMethodSet(T types.Type, msets *MethodSetCache) []*types.Selection {
	return typeutil.IntuitiveMethodSet(T, msets)
}

func MakeHasher() Hasher {
	return typeutil.MakeHasher()
}

type Map[V any] struct {
	m typeutil.Map
}

func (m *Map[V]) Delete(key types.Type) bool { return m.m.Delete(key) }
func (m *Map[V]) At(key types.Type) (V, bool) {
	v := m.m.At(key)
	if v == nil {
		var zero V
		return zero, false
	} else {
		return v.(V), true
	}
}
func (m *Map[V]) Set(key types.Type, value V) { m.m.Set(key, value) }
func (m *Map[V]) Len() int                    { return m.m.Len() }
func (m *Map[V]) Iterate(f func(key types.Type, value V)) {
	ff := func(key types.Type, value interface{}) {
		f(key, value.(V))
	}
	m.m.Iterate(ff)

}
func (m *Map[V]) Keys() []types.Type          { return m.m.Keys() }
func (m *Map[V]) String() string              { return m.m.String() }
func (m *Map[V]) KeysString() string          { return m.m.KeysString() }
func (m *Map[V]) SetHasher(h typeutil.Hasher) { m.m.SetHasher(h) }

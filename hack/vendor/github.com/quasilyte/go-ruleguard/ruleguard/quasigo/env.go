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

package quasigo

type funcKey struct {
	qualifier string
	name      string
}

func (k funcKey) String() string {
	if k.qualifier != "" {
		return k.qualifier + "." + k.name
	}
	return k.name
}

type nativeFunc struct {
	mappedFunc func(*ValueStack)
	name       string // Needed for the readable disasm
}

func newEnv() *Env {
	return &Env{
		nameToNativeFuncID: make(map[funcKey]uint16),
		nameToFuncID:       make(map[funcKey]uint16),

		debug: newDebugInfo(),
	}
}

func (env *Env) addNativeFunc(key funcKey, f func(*ValueStack)) {
	id := len(env.nativeFuncs)
	env.nativeFuncs = append(env.nativeFuncs, nativeFunc{
		mappedFunc: f,
		name:       key.String(),
	})
	env.nameToNativeFuncID[key] = uint16(id)
}

func (env *Env) addFunc(key funcKey, f *Func) {
	id := len(env.userFuncs)
	env.userFuncs = append(env.userFuncs, f)
	env.nameToFuncID[key] = uint16(id)
}

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

package qstrconv

import (
	"strconv"

	"github.com/quasilyte/go-ruleguard/ruleguard/quasigo"
)

func ImportAll(env *quasigo.Env) {
	env.AddNativeFunc(`strconv`, `Atoi`, Atoi)
	env.AddNativeFunc(`strconv`, `Itoa`, Itoa)
}

func Atoi(stack *quasigo.ValueStack) {
	s := stack.Pop().(string)
	v, err := strconv.Atoi(s)
	stack.PushInt(v)
	stack.Push(err)
}

func Itoa(stack *quasigo.ValueStack) {
	i := stack.PopInt()
	stack.Push(strconv.Itoa(i))
}

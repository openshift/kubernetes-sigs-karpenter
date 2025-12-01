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

package qstrings

import (
	"strings"

	"github.com/quasilyte/go-ruleguard/ruleguard/quasigo"
)

func ImportAll(env *quasigo.Env) {
	env.AddNativeFunc(`strings`, `Replace`, Replace)
	env.AddNativeFunc(`strings`, `ReplaceAll`, ReplaceAll)
	env.AddNativeFunc(`strings`, `TrimPrefix`, TrimPrefix)
	env.AddNativeFunc(`strings`, `TrimSuffix`, TrimSuffix)
	env.AddNativeFunc(`strings`, `HasPrefix`, HasPrefix)
	env.AddNativeFunc(`strings`, `HasSuffix`, HasSuffix)
	env.AddNativeFunc(`strings`, `Contains`, Contains)
}

func Replace(stack *quasigo.ValueStack) {
	n := stack.PopInt()
	newPart := stack.Pop().(string)
	oldPart := stack.Pop().(string)
	s := stack.Pop().(string)
	stack.Push(strings.Replace(s, oldPart, newPart, n))
}

func ReplaceAll(stack *quasigo.ValueStack) {
	newPart := stack.Pop().(string)
	oldPart := stack.Pop().(string)
	s := stack.Pop().(string)
	stack.Push(strings.ReplaceAll(s, oldPart, newPart))
}

func TrimPrefix(stack *quasigo.ValueStack) {
	prefix := stack.Pop().(string)
	s := stack.Pop().(string)
	stack.Push(strings.TrimPrefix(s, prefix))
}

func TrimSuffix(stack *quasigo.ValueStack) {
	prefix := stack.Pop().(string)
	s := stack.Pop().(string)
	stack.Push(strings.TrimSuffix(s, prefix))
}

func HasPrefix(stack *quasigo.ValueStack) {
	prefix := stack.Pop().(string)
	s := stack.Pop().(string)
	stack.Push(strings.HasPrefix(s, prefix))
}

func HasSuffix(stack *quasigo.ValueStack) {
	suffix := stack.Pop().(string)
	s := stack.Pop().(string)
	stack.Push(strings.HasSuffix(s, suffix))
}

func Contains(stack *quasigo.ValueStack) {
	substr := stack.Pop().(string)
	s := stack.Pop().(string)
	stack.Push(strings.Contains(s, substr))
}

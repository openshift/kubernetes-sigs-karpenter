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

package terminfo

type stack []interface{}

func (s *stack) push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack) pop() interface{} {
	if len(*s) == 0 {
		return nil
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *stack) popInt() int {
	if i, ok := s.pop().(int); ok {
		return i
	}
	return 0
}

func (s *stack) popBool() bool {
	if b, ok := s.pop().(bool); ok {
		return b
	}
	return false
}

func (s *stack) popByte() byte {
	if b, ok := s.pop().(byte); ok {
		return b
	}
	return 0
}

func (s *stack) popString() string {
	if a, ok := s.pop().(string); ok {
		return a
	}
	return ""
}

func (s *stack) reset() {
	*s = (*s)[:0]
}

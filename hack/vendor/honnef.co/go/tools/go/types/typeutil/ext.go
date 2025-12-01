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
	"fmt"
	"go/types"
)

type Iterator struct {
	elem types.Type
}

func (t *Iterator) Underlying() types.Type { return t }
func (t *Iterator) String() string         { return fmt.Sprintf("iterator(%s)", t.elem) }
func (t *Iterator) Elem() types.Type       { return t.elem }

func NewIterator(elem types.Type) *Iterator {
	return &Iterator{elem: elem}
}

type DeferStack struct{}

func (t *DeferStack) Underlying() types.Type { return t }
func (t *DeferStack) String() string         { return "deferStack" }

func NewDeferStack() *DeferStack {
	return &DeferStack{}
}

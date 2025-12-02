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

package dsl

import (
	"github.com/quasilyte/go-ruleguard/dsl/types"
)

// VarFilterContext carries Var and environment information into the filter function.
// It's an input parameter type for the Var.Filter function callback.
type VarFilterContext struct {
	// Type is mapped to Var.Type field.
	Type types.Type
}

// SizeOf returns the size of the given type.
// It uses the ruleguard.Context.Sizes to calculate the result.
func (*VarFilterContext) SizeOf(x types.Type) int { return 0 }

// GetType finds a type value by a given name.
// If a type can't be found (or a name is malformed), this function panics.
func (*VarFilterContext) GetType(name typeName) types.Type { return nil }

// GetInterface finds a type value that represents an interface by a given name.
// Works like `types.AsInterface(ctx.GetType(name))`.
func (*VarFilterContext) GetInterface(name typeName) *types.Interface { return nil }

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

package decorator

import (
	"go/ast"

	"github.com/dave/dst"
)

func newMap() Map {
	return Map{
		Ast: AstMap{
			Nodes:   map[dst.Node]ast.Node{},
			Scopes:  map[*dst.Scope]*ast.Scope{},
			Objects: map[*dst.Object]*ast.Object{},
		},
		Dst: DstMap{
			Nodes:   map[ast.Node]dst.Node{},
			Scopes:  map[*ast.Scope]*dst.Scope{},
			Objects: map[*ast.Object]*dst.Object{},
		},
	}
}

// Map holds a record of the mapping between ast and dst nodes, objects and scopes.
type Map struct {
	Ast AstMap
	Dst DstMap
}

// AstMap holds a record of the mapping from dst to ast nodes, objects and scopes.
type AstMap struct {
	Nodes   map[dst.Node]ast.Node       // Mapping from dst to ast Nodes
	Objects map[*dst.Object]*ast.Object // Mapping from dst to ast Objects
	Scopes  map[*dst.Scope]*ast.Scope   // Mapping from dst to ast Scopes
}

// DstMap holds a record of the mapping from ast to dst nodes, objects and scopes.
type DstMap struct {
	Nodes   map[ast.Node]dst.Node       // Mapping from ast to dst Nodes
	Objects map[*ast.Object]*dst.Object // Mapping from ast to dst Objects
	Scopes  map[*ast.Scope]*dst.Scope   // Mapping from ast to dst Scopes
}

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

// Package schutils provides tools to save or clone a schema
// when flattening a spec.
package schutils

import (
	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"
)

// Save registers a schema as an entry in spec #/definitions
func Save(sp *spec.Swagger, name string, schema *spec.Schema) {
	if schema == nil {
		return
	}

	if sp.Definitions == nil {
		sp.Definitions = make(map[string]spec.Schema, 150)
	}

	sp.Definitions[name] = *schema
}

// Clone deep-clones a schema
func Clone(schema *spec.Schema) *spec.Schema {
	var sch spec.Schema
	_ = swag.FromDynamicJSON(schema, &sch)

	return &sch
}

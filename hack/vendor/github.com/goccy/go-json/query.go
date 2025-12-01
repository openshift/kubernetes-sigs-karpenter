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

package json

import (
	"github.com/goccy/go-json/internal/encoder"
)

type (
	// FieldQuery you can dynamically filter the fields in the structure by creating a FieldQuery,
	// adding it to context.Context using SetFieldQueryToContext and then passing it to MarshalContext.
	// This is a type-safe operation, so it is faster than filtering using map[string]interface{}.
	FieldQuery       = encoder.FieldQuery
	FieldQueryString = encoder.FieldQueryString
)

var (
	// FieldQueryFromContext get current FieldQuery from context.Context.
	FieldQueryFromContext = encoder.FieldQueryFromContext
	// SetFieldQueryToContext set current FieldQuery to context.Context.
	SetFieldQueryToContext = encoder.SetFieldQueryToContext
)

// BuildFieldQuery builds FieldQuery by fieldName or sub field query.
// First, specify the field name that you want to keep in structure type.
// If the field you want to keep is a structure type, by creating a sub field query using BuildSubFieldQuery,
// you can select the fields you want to keep in the structure.
// This description can be written recursively.
func BuildFieldQuery(fields ...FieldQueryString) (*FieldQuery, error) {
	query, err := Marshal(fields)
	if err != nil {
		return nil, err
	}
	return FieldQueryString(query).Build()
}

// BuildSubFieldQuery builds sub field query.
func BuildSubFieldQuery(name string) *SubFieldQuery {
	return &SubFieldQuery{name: name}
}

type SubFieldQuery struct {
	name string
}

func (q *SubFieldQuery) Fields(fields ...FieldQueryString) FieldQueryString {
	query, _ := Marshal(map[string][]FieldQueryString{q.name: fields})
	return FieldQueryString(query)
}

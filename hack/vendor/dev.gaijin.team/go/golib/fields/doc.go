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

// Package fields provides types and functions to work with key-value pairs.
//
// The package offers three primary abstractions:
//
//   - Field: A key-value pair where the key is a string and the value can be any type.
//   - Dict: A map-based collection of unique fields, providing efficient key-based lookup.
//   - List: An ordered collection of fields that preserves insertion order.
//
// Fields can be created using the F constructor, and both Dict and List provide
// conversion methods between the two collection types. All types implement String()
// for consistent string representation.
//
// Example usage:
//
//	// Create fields
//	f1 := fields.F("status", "success")
//	f2 := fields.F("code", 200)
//
//	// Working with a List (ordered collection)
//	var list fields.List
//	list.Add(f1, f2)
//	fmt.Println(list) // "(status=success, code=200)"
//
//	// Working with a Dict (unique key collection)
//	dict := fields.Dict{}
//	dict.Add(f1, f2, fields.F("status", "updated")) // overwrites "status"
//	fmt.Println(dict) // "(status=updated, code=200)" (order may vary)
//
//	// Converting between types
//	list2 := dict.ToList() // order unspecified
//	dict2 := list.ToDict() // last occurrence of each key wins
package fields

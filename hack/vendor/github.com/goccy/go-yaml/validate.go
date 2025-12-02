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

package yaml

// StructValidator need to implement Struct method only
// ( see https://pkg.go.dev/github.com/go-playground/validator/v10#Validate.Struct )
type StructValidator interface {
	Struct(interface{}) error
}

// FieldError need to implement StructField method only
// ( see https://pkg.go.dev/github.com/go-playground/validator/v10#FieldError )
type FieldError interface {
	StructField() string
}

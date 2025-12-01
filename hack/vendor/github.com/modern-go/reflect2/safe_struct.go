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

package reflect2

type safeStructType struct {
	safeType
}

func (type2 *safeStructType) FieldByName(name string) StructField {
	field, found := type2.Type.FieldByName(name)
	if !found {
		panic("field " + name + " not found")
	}
	return &safeField{StructField: field}
}

func (type2 *safeStructType) Field(i int) StructField {
	return &safeField{StructField: type2.Type.Field(i)}
}

func (type2 *safeStructType) FieldByIndex(index []int) StructField {
	return &safeField{StructField: type2.Type.FieldByIndex(index)}
}

func (type2 *safeStructType) FieldByNameFunc(match func(string) bool) StructField {
	field, found := type2.Type.FieldByNameFunc(match)
	if !found {
		panic("field match condition not found in " + type2.Type.String())
	}
	return &safeField{StructField: field}
}

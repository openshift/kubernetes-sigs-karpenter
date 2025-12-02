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

// +build gofuzz

package syntax

// Fuzz is the input point for go-fuzz
func Fuzz(data []byte) int {
	sdata := string(data)
	tree, err := Parse(sdata, RegexOptions(0))
	if err != nil {
		return 0
	}

	// translate it to code
	_, err = Write(tree)
	if err != nil {
		panic(err)
	}

	return 1
}

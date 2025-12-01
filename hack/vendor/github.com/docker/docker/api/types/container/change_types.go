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

package container

const (
	// ChangeModify represents the modify operation.
	ChangeModify ChangeType = 0
	// ChangeAdd represents the add operation.
	ChangeAdd ChangeType = 1
	// ChangeDelete represents the delete operation.
	ChangeDelete ChangeType = 2
)

func (ct ChangeType) String() string {
	switch ct {
	case ChangeModify:
		return "C"
	case ChangeAdd:
		return "A"
	case ChangeDelete:
		return "D"
	default:
		return ""
	}
}

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

package specificity

type StandardMatch struct{}

func (s StandardMatch) IsMoreSpecific(than MatchSpecificity) bool {
	return isMoreSpecific(s, than)
}

func (s StandardMatch) Equal(to MatchSpecificity) bool {
	return equalSpecificity(s, to)
}

func (s StandardMatch) class() specificityClass {
	return StandardClass
}

func (s StandardMatch) String() string {
	return "Standard"
}

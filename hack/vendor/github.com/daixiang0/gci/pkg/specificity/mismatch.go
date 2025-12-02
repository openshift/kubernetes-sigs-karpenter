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

type MisMatch struct{}

func (m MisMatch) IsMoreSpecific(than MatchSpecificity) bool {
	return isMoreSpecific(m, than)
}

func (m MisMatch) Equal(to MatchSpecificity) bool {
	return equalSpecificity(m, to)
}

func (m MisMatch) class() specificityClass {
	return MisMatchClass
}

func (m MisMatch) String() string {
	return "Mismatch"
}

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

type specificityClass int

const (
	MisMatchClass    = 0
	DefaultClass     = 10
	StandardClass    = 20
	MatchClass       = 30
	NameClass        = 40
	LocalModuleClass = 50
)

// MatchSpecificity is used to determine which section matches an import best
type MatchSpecificity interface {
	IsMoreSpecific(than MatchSpecificity) bool
	Equal(to MatchSpecificity) bool
	class() specificityClass
}

func isMoreSpecific(this, than MatchSpecificity) bool {
	return this.class() > than.class()
}

func equalSpecificity(base, to MatchSpecificity) bool {
	// m.class() == to.class() would not work for Match
	return !base.IsMoreSpecific(to) && !to.IsMoreSpecific(base)
}

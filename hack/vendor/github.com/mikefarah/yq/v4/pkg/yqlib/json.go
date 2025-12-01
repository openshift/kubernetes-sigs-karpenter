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

package yqlib

type JsonPreferences struct {
	Indent        int
	ColorsEnabled bool
	UnwrapScalar  bool
}

func NewDefaultJsonPreferences() JsonPreferences {
	return JsonPreferences{
		Indent:        2,
		ColorsEnabled: true,
		UnwrapScalar:  true,
	}
}

func (p *JsonPreferences) Copy() JsonPreferences {
	return JsonPreferences{
		Indent:        p.Indent,
		ColorsEnabled: p.ColorsEnabled,
		UnwrapScalar:  p.UnwrapScalar,
	}
}

var ConfiguredJSONPreferences = NewDefaultJsonPreferences()

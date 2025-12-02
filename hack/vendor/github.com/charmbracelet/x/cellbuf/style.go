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

package cellbuf

import (
	"github.com/charmbracelet/colorprofile"
)

// Convert converts a style to respect the given color profile.
func ConvertStyle(s Style, p colorprofile.Profile) Style {
	switch p {
	case colorprofile.TrueColor:
		return s
	case colorprofile.Ascii:
		s.Fg = nil
		s.Bg = nil
		s.Ul = nil
	case colorprofile.NoTTY:
		return Style{}
	}

	if s.Fg != nil {
		s.Fg = p.Convert(s.Fg)
	}
	if s.Bg != nil {
		s.Bg = p.Convert(s.Bg)
	}
	if s.Ul != nil {
		s.Ul = p.Convert(s.Ul)
	}

	return s
}

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

package termenv

import (
	"strings"

	"github.com/aymanbagabas/go-osc52/v2"
)

// Copy copies text to clipboard using OSC 52 escape sequence.
func (o Output) Copy(str string) {
	s := osc52.New(str)
	if strings.HasPrefix(o.environ.Getenv("TERM"), "screen") {
		s = s.Screen()
	}
	_, _ = s.WriteTo(o)
}

// CopyPrimary copies text to primary clipboard (X11) using OSC 52 escape
// sequence.
func (o Output) CopyPrimary(str string) {
	s := osc52.New(str).Primary()
	if strings.HasPrefix(o.environ.Getenv("TERM"), "screen") {
		s = s.Screen()
	}
	_, _ = s.WriteTo(o)
}

// Copy copies text to clipboard using OSC 52 escape sequence.
func Copy(str string) {
	output.Copy(str)
}

// CopyPrimary copies text to primary clipboard (X11) using OSC 52 escape
// sequence.
func CopyPrimary(str string) {
	output.CopyPrimary(str)
}

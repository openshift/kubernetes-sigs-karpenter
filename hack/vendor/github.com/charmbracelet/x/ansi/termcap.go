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

package ansi

import (
	"encoding/hex"
	"strings"
)

// RequestTermcap (XTGETTCAP) requests Termcap/Terminfo strings.
//
//	DCS + q <Pt> ST
//
// Where <Pt> is a list of Termcap/Terminfo capabilities, encoded in 2-digit
// hexadecimals, separated by semicolons.
//
// See: https://man7.org/linux/man-pages/man5/terminfo.5.html
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h3-Operating-System-Commands
func XTGETTCAP(caps ...string) string {
	if len(caps) == 0 {
		return ""
	}

	s := "\x1bP+q"
	for i, c := range caps {
		if i > 0 {
			s += ";"
		}
		s += strings.ToUpper(hex.EncodeToString([]byte(c)))
	}

	return s + "\x1b\\"
}

// RequestTermcap is an alias for [XTGETTCAP].
func RequestTermcap(caps ...string) string {
	return XTGETTCAP(caps...)
}

// RequestTerminfo is an alias for [XTGETTCAP].
func RequestTerminfo(caps ...string) string {
	return XTGETTCAP(caps...)
}

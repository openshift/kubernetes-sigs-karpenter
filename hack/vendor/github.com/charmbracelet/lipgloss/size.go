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

package lipgloss

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
)

// Width returns the cell width of characters in the string. ANSI sequences are
// ignored and characters wider than one cell (such as Chinese characters and
// emojis) are appropriately measured.
//
// You should use this instead of len(string) len([]rune(string) as neither
// will give you accurate results.
func Width(str string) (width int) {
	for _, l := range strings.Split(str, "\n") {
		w := ansi.StringWidth(l)
		if w > width {
			width = w
		}
	}

	return width
}

// Height returns height of a string in cells. This is done simply by
// counting \n characters. If your strings use \r\n for newlines you should
// convert them to \n first, or simply write a separate function for measuring
// height.
func Height(str string) int {
	return strings.Count(str, "\n") + 1
}

// Size returns the width and height of the string in cells. ANSI sequences are
// ignored and characters wider than one cell (such as Chinese characters and
// emojis) are appropriately measured.
func Size(str string) (width, height int) {
	width = Width(str)
	height = Height(str)
	return width, height
}

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

package parser

import "fmt"

const (
	colorFgHiBlack int = iota + 90
	colorFgHiRed
	colorFgHiGreen
	colorFgHiYellow
	colorFgHiBlue
	colorFgHiMagenta
	colorFgHiCyan
)

var colorTable = []int{
	colorFgHiRed,
	colorFgHiGreen,
	colorFgHiYellow,
	colorFgHiBlue,
	colorFgHiMagenta,
	colorFgHiCyan,
}

func colorize(idx int, content string) string {
	colorIdx := idx % len(colorTable)
	color := colorTable[colorIdx]
	return fmt.Sprintf("\x1b[1;%dm", color) + content + "\x1b[22;0m"
}

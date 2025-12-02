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

// This source inspired by https://github.com/fatih/color.
package printer

import (
	"fmt"
	"strings"
)

type ColorAttribute int

const (
	ColorReset ColorAttribute = iota
	ColorBold
	ColorFaint
	ColorItalic
	ColorUnderline
	ColorBlinkSlow
	ColorBlinkRapid
	ColorReverseVideo
	ColorConcealed
	ColorCrossedOut
)

const (
	ColorFgHiBlack ColorAttribute = iota + 90
	ColorFgHiRed
	ColorFgHiGreen
	ColorFgHiYellow
	ColorFgHiBlue
	ColorFgHiMagenta
	ColorFgHiCyan
	ColorFgHiWhite
)

const (
	ColorResetBold ColorAttribute = iota + 22
	ColorResetItalic
	ColorResetUnderline
	ColorResetBlinking

	ColorResetReversed
	ColorResetConcealed
	ColorResetCrossedOut
)

const escape = "\x1b"

var colorResetMap = map[ColorAttribute]ColorAttribute{
	ColorBold:         ColorResetBold,
	ColorFaint:        ColorResetBold,
	ColorItalic:       ColorResetItalic,
	ColorUnderline:    ColorResetUnderline,
	ColorBlinkSlow:    ColorResetBlinking,
	ColorBlinkRapid:   ColorResetBlinking,
	ColorReverseVideo: ColorResetReversed,
	ColorConcealed:    ColorResetConcealed,
	ColorCrossedOut:   ColorResetCrossedOut,
}

func format(attrs ...ColorAttribute) string {
	format := make([]string, 0, len(attrs))
	for _, attr := range attrs {
		format = append(format, fmt.Sprint(attr))
	}
	return fmt.Sprintf("%s[%sm", escape, strings.Join(format, ";"))
}

func unformat(attrs ...ColorAttribute) string {
	format := make([]string, len(attrs))
	for _, attr := range attrs {
		v := fmt.Sprint(ColorReset)
		reset, exists := colorResetMap[attr]
		if exists {
			v = fmt.Sprint(reset)
		}
		format = append(format, v)
	}
	return fmt.Sprintf("%s[%sm", escape, strings.Join(format, ";"))
}

func colorize(msg string, attrs ...ColorAttribute) string {
	return format(attrs...) + msg + unformat(attrs...)
}

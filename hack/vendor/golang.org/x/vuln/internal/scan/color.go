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

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package scan

const (
	// These are all the constants for the terminal escape strings

	colorEscape = "\033["
	colorEnd    = "m"

	colorReset     = colorEscape + "0" + colorEnd
	colorBold      = colorEscape + "1" + colorEnd
	colorFaint     = colorEscape + "2" + colorEnd
	colorUnderline = colorEscape + "4" + colorEnd
	colorBlink     = colorEscape + "5" + colorEnd

	fgBlack   = colorEscape + "30" + colorEnd
	fgRed     = colorEscape + "31" + colorEnd
	fgGreen   = colorEscape + "32" + colorEnd
	fgYellow  = colorEscape + "33" + colorEnd
	fgBlue    = colorEscape + "34" + colorEnd
	fgMagenta = colorEscape + "35" + colorEnd
	fgCyan    = colorEscape + "36" + colorEnd
	fgWhite   = colorEscape + "37" + colorEnd

	bgBlack   = colorEscape + "40" + colorEnd
	bgRed     = colorEscape + "41" + colorEnd
	bgGreen   = colorEscape + "42" + colorEnd
	bgYellow  = colorEscape + "43" + colorEnd
	bgBlue    = colorEscape + "44" + colorEnd
	bgMagenta = colorEscape + "45" + colorEnd
	bgCyan    = colorEscape + "46" + colorEnd
	bgWhite   = colorEscape + "47" + colorEnd

	fgBlackHi   = colorEscape + "90" + colorEnd
	fgRedHi     = colorEscape + "91" + colorEnd
	fgGreenHi   = colorEscape + "92" + colorEnd
	fgYellowHi  = colorEscape + "93" + colorEnd
	fgBlueHi    = colorEscape + "94" + colorEnd
	fgMagentaHi = colorEscape + "95" + colorEnd
	fgCyanHi    = colorEscape + "96" + colorEnd
	fgWhiteHi   = colorEscape + "97" + colorEnd

	bgBlackHi   = colorEscape + "100" + colorEnd
	bgRedHi     = colorEscape + "101" + colorEnd
	bgGreenHi   = colorEscape + "102" + colorEnd
	bgYellowHi  = colorEscape + "103" + colorEnd
	bgBlueHi    = colorEscape + "104" + colorEnd
	bgMagentaHi = colorEscape + "105" + colorEnd
	bgCyanHi    = colorEscape + "106" + colorEnd
	bgWhiteHi   = colorEscape + "107" + colorEnd
)

const (
	_ = colorReset
	_ = colorBold
	_ = colorFaint
	_ = colorUnderline
	_ = colorBlink

	_ = fgBlack
	_ = fgRed
	_ = fgGreen
	_ = fgYellow
	_ = fgBlue
	_ = fgMagenta
	_ = fgCyan
	_ = fgWhite

	_ = fgBlackHi
	_ = fgRedHi
	_ = fgGreenHi
	_ = fgYellowHi
	_ = fgBlueHi
	_ = fgMagentaHi
	_ = fgCyanHi
	_ = fgWhiteHi

	_ = bgBlack
	_ = bgRed
	_ = bgGreen
	_ = bgYellow
	_ = bgBlue
	_ = bgMagenta
	_ = bgCyan
	_ = bgWhite

	_ = bgBlackHi
	_ = bgRedHi
	_ = bgGreenHi
	_ = bgYellowHi
	_ = bgBlueHi
	_ = bgMagentaHi
	_ = bgCyanHi
	_ = bgWhiteHi
)

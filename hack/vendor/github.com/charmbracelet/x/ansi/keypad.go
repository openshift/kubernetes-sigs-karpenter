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

// Keypad Application Mode (DECKPAM) is a mode that determines whether the
// keypad sends application sequences or ANSI sequences.
//
// This works like enabling [DECNKM].
// Use [NumericKeypadMode] to set the numeric keypad mode.
//
//	ESC =
//
// See: https://vt100.net/docs/vt510-rm/DECKPAM.html
const (
	KeypadApplicationMode = "\x1b="
	DECKPAM               = KeypadApplicationMode
)

// Keypad Numeric Mode (DECKPNM) is a mode that determines whether the keypad
// sends application sequences or ANSI sequences.
//
// This works the same as disabling [DECNKM].
//
//	ESC >
//
// See: https://vt100.net/docs/vt510-rm/DECKPNM.html
const (
	KeypadNumericMode = "\x1b>"
	DECKPNM           = KeypadNumericMode
)

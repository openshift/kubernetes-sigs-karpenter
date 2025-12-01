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

// SetIconNameWindowTitle returns a sequence for setting the icon name and
// window title.
//
//	OSC 0 ; title ST
//	OSC 0 ; title BEL
//
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h2-Operating-System-Commands
func SetIconNameWindowTitle(s string) string {
	return "\x1b]0;" + s + "\x07"
}

// SetIconName returns a sequence for setting the icon name.
//
//	OSC 1 ; title ST
//	OSC 1 ; title BEL
//
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h2-Operating-System-Commands
func SetIconName(s string) string {
	return "\x1b]1;" + s + "\x07"
}

// SetWindowTitle returns a sequence for setting the window title.
//
//	OSC 2 ; title ST
//	OSC 2 ; title BEL
//
// See: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h2-Operating-System-Commands
func SetWindowTitle(s string) string {
	return "\x1b]2;" + s + "\x07"
}

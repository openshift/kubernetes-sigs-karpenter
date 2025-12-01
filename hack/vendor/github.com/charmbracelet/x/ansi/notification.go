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

// Notify sends a desktop notification using iTerm's OSC 9.
//
//	OSC 9 ; Mc ST
//	OSC 9 ; Mc BEL
//
// Where Mc is the notification body.
//
// See: https://iterm2.com/documentation-escape-codes.html
func Notify(s string) string {
	return "\x1b]9;" + s + "\x07"
}

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

import "fmt"

// ITerm2 returns a sequence that uses the iTerm2 proprietary protocol. Use the
// iterm2 package for a more convenient API.
//
//	OSC 1337 ; key = value ST
//
// Example:
//
//	ITerm2(iterm2.File{...})
//
// See https://iterm2.com/documentation-escape-codes.html
// See https://iterm2.com/documentation-images.html
func ITerm2(data any) string {
	return "\x1b]1337;" + fmt.Sprint(data) + "\x07"
}
